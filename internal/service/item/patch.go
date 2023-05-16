package item

import (
	"context"

	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

func (s *Service) Patch(ctx context.Context, req *desc.PatchRequest) (*item.Item, error) {
	res, err := s.itemRepository.Patch(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
