package item

import (
	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
)

type Service struct {
	itemRepository item.Repository
}

func NewService(itemRepository item.Repository) *Service {
	return &Service{
		itemRepository: itemRepository,
	}
}
