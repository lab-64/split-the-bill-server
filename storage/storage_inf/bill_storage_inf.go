package storage_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type IBillStorage interface {
	Create(bill BillModel) error

	GetByID(id UUID) (BillModel, error)
}
