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
}
