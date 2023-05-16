package item_v1

import (
	"context"

	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Get(ctx context.Context, e *emptypb.Empty) (*desc.GetResponse, error) {
	res, err := i.itemService.Get(ctx, e)
	if err != nil {
		return nil, err
	}
	var resDesc []*desc.Item
	for _, item := range res {
		resDesc = append(resDesc, &desc.Item{
			Id:         item.Id,
			CampaignId: item.CampaignId,
			Item: &desc.ItemInfo{
				Name:        item.Item.Name,
				Description: item.Item.Description,
			},
			Priority: item.Priority,
			Removed:  item.Removed,
		})
	}
	return &desc.GetResponse{
		Items: resDesc,
	}, nil
}
