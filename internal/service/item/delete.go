package item

import (
	"context"

	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

// Delete service
func (s *Service) Delete(ctx context.Context, req *desc.DeleteRequest) (*item.DeleteInfo, error) {
	res, err := s.itemRepository.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
