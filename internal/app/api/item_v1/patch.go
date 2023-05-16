package item_v1

import (
	"context"

	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

func (i *Implementation) Patch(ctx context.Context, req *desc.PatchRequest) (*desc.PatchResponse, error) {
	item, err := i.itemService.Patch(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.PatchResponse{
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
