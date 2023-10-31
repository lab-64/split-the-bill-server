package storage_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type IBillStorage interface {
	Create(bill model.Bill) error

	GetByID(id uuid.UUID) (model.Bill, error)
}
