package storage

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type IBillStorage interface {
	// Create creates a new bill and returns the created bill
	Create(bill model.Bill) (model.Bill, error)

	// UpdateBill updates the stored bill with the given bill
	UpdateBill(bill model.Bill) (model.Bill, error)

	// GetByID returns a bill by its id
	GetByID(id UUID) (model.Bill, error)

	// DeleteBill deletes the bill with the given id
	DeleteBill(id UUID) error

	// CreateItem creates an item for a bill
	CreateItem(item model.Item) (model.Item, error)

	// GetItemByID returns an item by its id
	GetItemByID(id UUID) (model.Item, error)

	// UpdateItem updates the stored item with the given item
	UpdateItem(item model.Item) (model.Item, error)

	// DeleteItem deletes the item with the given id
	DeleteItem(id UUID) error
}
