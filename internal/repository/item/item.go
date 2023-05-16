package item

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	tableName = "items"
	table     = "campaigns"
)

type Repository interface {
	Post(ctx context.Context, req *desc.PostRequest) (*Item, error)
	Get(ctx context.Context, _ *emptypb.Empty) ([]*Item, error)
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
	CreatedAt  timestamp.Timestamp
	UpdatedAt  timestamp.Timestamp
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
	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns("name").
		Values(req.GetInfo().GetName()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	builderItems := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("campaign_id", "name", "priority").
		Values(id, req.GetInfo().GetName(), priority+1)

	query, args, err = builderItems.ToSql()
	if err != nil {
		return nil, err
	}

	builderRes := sq.Select("id", "campaign_id", "name", "description", "priority", "removed", "createdAt").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err = builderRes.ToSql()
	if err != nil {
		return nil, err
	}

	row, err = r.db.QueryContext(ctx, query, args...)
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

func (r *repository) Get(ctx context.Context) ([]*Item, error) {
	builder := sq.Select("id", "campaign_id", "name", "description", "priority", "removed", "createdAt").
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

	return resDesc, nil
}

func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) (*DeleteInfo, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId(), "campaign_id": req.GetCampaignId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
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
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("name", req.GetUpdateInfo().Name).
		Set("description", req.GetUpdateInfo().Description).
		Where(sq.Eq{"id": req.GetId(), "campaign_id": req.GetCampaignId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var item *Item
	err = row.Scan(&item.Id, &item.CampaignId, &item.Item.Name, &item.Item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
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
