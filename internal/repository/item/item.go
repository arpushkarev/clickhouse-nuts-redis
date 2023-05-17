package item

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/redis"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	tableName = "items"
	table     = "campaigns"
)

type Repository interface {
	Post(ctx context.Context, req *desc.PostRequest) (*Item, error)
	Get(ctx context.Context, e *emptypb.Empty) ([]*Item, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) (*DeleteInfo, error)
	Patch(ctx context.Context, req *desc.PatchRequest) (*Item, error)
}

type repository struct {
	db *sqlx.DB
}

type Info struct {
	Name        string
	Description string
}

type Item struct {
	Id         int64
	CampaignId int64
	Item       *Info
	Priority   int64
	Removed    bool
	CreatedAt  timestamppb.Timestamp
	UpdatedAt  timestamppb.Timestamp
}

type DeleteInfo struct {
	Id         int64
	CampaignId int64
	Removed    bool
}

var priority int64 = 1

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Post(ctx context.Context, req *desc.PostRequest) (*Item, error) {
	name := req.GetInfo().GetName()

	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(name).
		Suffix("returning id")

	query1, args1, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row1, err := r.db.QueryContext(ctx, query1, args1...)
	if err != nil {
		return nil, err
	}
	defer row1.Close()

	row1.Next()
	var id int64
	err = row1.Scan(&id)
	if err != nil {
		return nil, err
	}

	builderItems := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("campaign_id", "name", "priority").
		Values(id, name, priority+1)

	query2, args2, err := builderItems.ToSql()
	if err != nil {
		return nil, err
	}

	row2, err := r.db.QueryContext(ctx, query2, args2...)
	if err != nil {
		return nil, err
	}
	defer row2.Close()

	builderRes := sq.Select("id", "campaign_id", "name", "description", "priority", "removed", "created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builderRes.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var res *Item
	err = row.Scan(&res.Id, &res.CampaignId, &res.Item.Name, &res.Item.Description, &res.Priority, &res.Removed, &res.CreatedAt)
	if err != nil {
		return nil, err
	}

	priority += 1

	return res, err
}

func (r *repository) Get(ctx context.Context, e *emptypb.Empty) ([]*Item, error) {

	cache, err := redis.New()
	if err != nil {
		log.Printf("Failed to create cache: %s", err.Error())
	}

	data, err := cache.Get(ctx, "items").Bytes()
	if err != nil {
		return nil, err
	}

	if data != nil {
		var item []*Item

		err = json.Unmarshal(data, &item)

		return item, nil
	}

	builder := sq.Select("id", "campaign_id", "name", "description", "priority", "removed", "created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var res []Item
	for row.Next() {
		var item Item
		err = row.Scan(&item.Id, &item.CampaignId, &item.Item.Name, &item.Item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, Item{
			Id:         item.Id,
			CampaignId: item.CampaignId,
			Item:       item.Item,
			Priority:   item.Priority,
			Removed:    item.Removed,
			CreatedAt:  item.CreatedAt,
		})
	}

	var resDesc []*Item
	for _, elem := range res {
		resDesc = append(resDesc, &Item{
			Id:         elem.Id,
			CampaignId: elem.CampaignId,
			Item:       elem.Item,
			Priority:   elem.Priority,
			Removed:    elem.Removed,
			CreatedAt:  elem.CreatedAt,
		})
	}

	p, err := json.Marshal(resDesc)

	err = cache.Set(ctx, "items", p, cache.Duration).Err()
	if err != nil {
		log.Printf("failed to set data into cahce:%s", err.Error())
	}

	return resDesc, nil
}

func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) (*DeleteInfo, error) {

	cache, err := redis.New()
	if err != nil {
		log.Printf("Failed to create cache: %s", err.Error())
	}

	err = cache.Del(ctx, "items").Err()

	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId(), "campaign_id": req.GetCampaignId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("errors.item.notFound")
		}
		return nil, err
	}
	defer row.Close()

	row.Next()
	var res *DeleteInfo
	err = row.Scan(&res.Id, &res.CampaignId, &res.Removed)
	if err != nil {
		return nil, err
	}

	return &DeleteInfo{
		Id:         res.Id,
		CampaignId: res.CampaignId,
		Removed:    res.Removed,
	}, nil
}

func (r *repository) Patch(ctx context.Context, req *desc.PatchRequest) (*Item, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("name", req.GetUpdateInfo().Name).
		Set("description", req.GetUpdateInfo().Description).
		Where(sq.Eq{"id": req.GetId(), "campaign_id": req.GetCampaignId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, query, args...)

	var item *Item
	err = row.Scan(&item.Id, &item.CampaignId, &item.Item.Name, &item.Item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Rollback()
			return nil, errors.New("errors.item.notFound")
		}

		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &Item{
		Id:         item.Id,
		CampaignId: item.CampaignId,
		Item:       item.Item,
		Priority:   item.Priority,
		Removed:    item.Removed,
		CreatedAt:  item.CreatedAt,
	}, nil
}
