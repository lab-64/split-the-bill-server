package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type BillStorage struct {
	e *Ephemeral
}

func NewBillStorage(ephemeral *Ephemeral) storage.IBillStorage {
	return &BillStorage{e: ephemeral}
}

func (b BillStorage) Create(bill types.Bill) error {
	b.e.lock.Lock()
	defer b.e.lock.Unlock()
	_, exists := b.e.bills[bill.ID]
	if exists {
		return storage.BillAlreadyExistsError
	}
	b.e.bills[bill.ID] = bill
	return nil
}

func (b BillStorage) GetByID(id uuid.UUID) (types.Bill, error) {
	b.e.lock.Lock()
	defer b.e.lock.Unlock()
	bill, exists := b.e.bills[id]
	if !exists {
		return bill, storage.NoSuchBillError
	}
	return bill, nil
}
