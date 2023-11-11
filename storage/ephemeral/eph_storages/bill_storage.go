package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/storage_inf"
)

type BillStorage struct {
	e *ephemeral.Ephemeral
}

func NewBillStorage(ephemeral *ephemeral.Ephemeral) storage_inf.IBillStorage {
	return &BillStorage{e: ephemeral}
}

func (b BillStorage) Create(bill model.BillModel) error {
	b.e.Lock.Lock()
	defer b.e.Lock.Unlock()
	_, exists := b.e.Bills[bill.ID]
	if exists {
		return storage.BillAlreadyExistsError
	}
	b.e.Bills[bill.ID] = &bill
	return nil
}

func (b BillStorage) GetByID(id uuid.UUID) (model.BillModel, error) {
	b.e.Lock.Lock()
	defer b.e.Lock.Unlock()
	bill, exists := b.e.Bills[id]
	if !exists {
		return *bill, storage.NoSuchBillError
	}
	return *bill, nil
}
