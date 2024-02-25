package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	eph "split-the-bill-server/storage/ephemeral"
)

type BillStorage struct {
	e *eph.Ephemeral
}

func NewBillStorage(ephemeral *eph.Ephemeral) storage.IBillStorage {
	return &BillStorage{e: ephemeral}
}

func (b BillStorage) UpdateBill(bill model.Bill) (model.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) Create(bill model.Bill) (model.Bill, error) {
	r := b.e.Locker.Lock(eph.RBills)
	defer b.e.Locker.Unlock(r)
	_, exists := b.e.Bills[bill.ID]
	if exists {
		return model.Bill{}, storage.BillAlreadyExistsError
	}
	b.e.Bills[bill.ID] = &bill
	return bill, nil
}

func (b BillStorage) GetByID(id uuid.UUID) (model.Bill, error) {
	r := b.e.Locker.Lock(eph.RBills)
	defer b.e.Locker.Unlock(r)
	bill, exists := b.e.Bills[id]
	if !exists {
		return *bill, storage.NoSuchBillError
	}
	return *bill, nil
}

func (b BillStorage) CreateItem(item model.Item) (model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) GetItemByID(id uuid.UUID) (model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) GetAllByUserID(userID uuid.UUID) ([]model.BillModel, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) UpdateItem(item model.Item) (model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (b BillStorage) DeleteItem(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
