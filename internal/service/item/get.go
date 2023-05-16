package item

import (
	"context"

	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) Get(ctx context.Context, e *emptypb.Empty) ([]*item.Item, error) {
	res, err := s.itemRepository.Get(ctx, e)
	if err != nil {
		return nil, err
	}

	return res, nil
}
