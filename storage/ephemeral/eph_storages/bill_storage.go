package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	eph "split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/storage_inf"
)

type BillStorage struct {
	e *eph.Ephemeral
}

func NewBillStorage(ephemeral *eph.Ephemeral) storage_inf.IBillStorage {
	return &BillStorage{e: ephemeral}
}

func (b BillStorage) Create(bill model.BillModel) (model.BillModel, error) {
	r := b.e.Locker.Lock(eph.RBills)
	defer b.e.Locker.Unlock(r)
	_, exists := b.e.Bills[bill.ID]
	if exists {
		return model.BillModel{}, storage.BillAlreadyExistsError
	}
	b.e.Bills[bill.ID] = &bill
	return bill, nil
}

func (b BillStorage) GetByID(id uuid.UUID) (model.BillModel, error) {
	r := b.e.Locker.Lock(eph.RBills)
	defer b.e.Locker.Unlock(r)
	bill, exists := b.e.Bills[id]
	if !exists {
		return *bill, storage.NoSuchBillError
	}
	return *bill, nil
}

func (b BillStorage) CreateItem(item model.ItemModel) (model.ItemModel, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) GetItemByID(id uuid.UUID) (model.ItemModel, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) UpdateItem(item model.ItemModel) (model.ItemModel, error) {
	//TODO implement me
	panic("implement me")
}
