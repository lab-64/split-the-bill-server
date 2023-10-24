package storage

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type IBillStorage interface {
	Create(bill types.Bill) error

	GetByID(id uuid.UUID) (types.Bill, error)
}
