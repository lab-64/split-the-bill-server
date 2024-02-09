package storage

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type IBillStorage interface {
	Create(bill model.Bill) (model.Bill, error)

	UpdateBill(bill model.Bill) (model.Bill, error)

	GetByID(id UUID) (model.Bill, error)

	// CreateItem creates an item for a bill
	CreateItem(item model.Item) (model.Item, error)

	// GetItemByID returns an item by its id
	GetItemByID(id UUID) (model.Item, error)

	// UpdateItem updates the stored item with the given item
	UpdateItem(item model.Item) (model.Item, error)
}
