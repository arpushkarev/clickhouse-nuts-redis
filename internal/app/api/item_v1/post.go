package item_v1

import (
	"context"

	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

func (i *Implementation) Post(ctx context.Context, req *desc.PostRequest) (*desc.PostResponse, error) {
	item, err := i.itemService.Post(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.PostResponse{
		Item: &desc.Item{
			Id:         item.Id,
			CampaignId: item.CampaignId,
			Item: &desc.ItemInfo{
				Name:        item.Item.Name,
				Description: item.Item.Description,
			},
			Priority: item.Priority,
			Removed:  item.Removed,
			//CreatedAt:  item.CreatedAt, - нужно проверить тип
		},
	}, nil
}
