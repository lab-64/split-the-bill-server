package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockBillCreate         func(model.BillModel) (model.BillModel, error)
	MockBillUpdate         func(model.BillModel) (model.BillModel, error)
	MockBillGetByID        func(uuid.UUID) (model.BillModel, error)
	MockBillCreateItem     func(model.ItemModel) (model.ItemModel, error)
	MockBillGetItemByID    func(uuid.UUID) (model.ItemModel, error)
	MockBillUpdateItem     func(model.ItemModel) (model.ItemModel, error)
	MockBillGetAllByUserID func(uuid.UUID) ([]model.BillModel, error)
)

func NewBillStorageMock() storage.IBillStorage {
	return &BillStorageMock{}
}

type BillStorageMock struct {
}

func (b BillStorageMock) Create(bill model.BillModel) (model.BillModel, error) {
	return MockBillCreate(bill)
}

func (b BillStorageMock) UpdateBill(bill model.BillModel) (model.BillModel, error) {
	return MockBillUpdate(bill)
}

func (b BillStorageMock) GetByID(id uuid.UUID) (model.BillModel, error) {
	return MockBillGetByID(id)
}

func (b BillStorageMock) GetAllByUserID(userID uuid.UUID) ([]model.BillModel, error) {
	return MockBillGetAllByUserID(userID)
}

func (b BillStorageMock) CreateItem(item model.ItemModel) (model.ItemModel, error) {
	return MockBillCreateItem(item)
}

func (b BillStorageMock) GetItemByID(id uuid.UUID) (model.ItemModel, error) {
	return MockBillGetItemByID(id)
}

func (b BillStorageMock) UpdateItem(item model.ItemModel) (model.ItemModel, error) {
	return MockBillUpdateItem(item)
}
