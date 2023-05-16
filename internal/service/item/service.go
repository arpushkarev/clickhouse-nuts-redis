package item

import (
	"github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
)

// Service structure
type Service struct {
	itemRepository item.Repository
}

// NewService initialisation
func NewService(itemRepository item.Repository) *Service {
	return &Service{
		itemRepository: itemRepository,
	}
}
