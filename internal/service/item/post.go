package item

import (
	"context"

	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

// Create service
func (s *Service) Post(ctx context.Context, req *desc.PostRequest) (*item.Item, error) {
	id, err := s.itemRepository.Post(ctx, req)
	if err != nil {
		return nil, err
	}

	return id, err
}
