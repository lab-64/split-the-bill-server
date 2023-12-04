package storage_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type IBillStorage interface {
	Create(bill BillModel) (BillModel, error)

	GetByID(id UUID) (BillModel, error)

	// CreateItem creates an item for a bill
	CreateItem(item ItemModel) (ItemModel, error)

	// GetItemByID returns an item by its id
	GetItemByID(id UUID) (ItemModel, error)

	// UpdateItem updates the stored item with the given item
	UpdateItem(item ItemModel) (ItemModel, error)
}
