package item_v1

import (
	"context"

	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

// Delete note by ID
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	item, err := i.itemService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.DeleteResponse{
		Id:         item.Id,
		CampaignId: item.CampaignId,
		Removed:    item.Removed,
	}, nil
}
