package item_v1

import (
	"github.com/arpushkarev/clickhouse-nuts-redis/internal/service/item"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
)

// Implementation new server
type Implementation struct {
	desc.UnimplementedItemV1Server

	itemService *item.Service
}

// NewImplementation starts connect
func NewImplementation(itemService *item.Service) *Implementation {
	return &Implementation{
		itemService: itemService,
	}
}
