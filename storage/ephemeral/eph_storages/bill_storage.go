package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
)

type BillStorage struct {
	e *ephemeral.Ephemeral
}

func NewBillStorage(ephemeral *ephemeral.Ephemeral) storage.IBillStorage {
	return &BillStorage{e: ephemeral}
}

func (b BillStorage) UpdateBill(bill model.BillModel) (model.BillModel, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) Create(bill model.BillModel) (model.BillModel, error) {
	b.e.Lock.Lock()
	defer b.e.Lock.Unlock()
	_, exists := b.e.Bills[bill.ID]
	if exists {
		return model.BillModel{}, storage.BillAlreadyExistsError
	}
	b.e.Bills[bill.ID] = &bill
	return bill, nil
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
