package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockBillCreate      func(model.Bill) (model.Bill, error)
	MockBillUpdate      func(model.Bill) (model.Bill, error)
	MockBillGetByID     func(uuid.UUID) (model.Bill, error)
	MockBillDelete      func(uuid.UUID) error
	MockBillCreateItem  func(model.Item) (model.Item, error)
	MockBillGetItemByID func(uuid.UUID) (model.Item, error)
	MockBillUpdateItem  func(model.Item) (model.Item, error)
	MockBillDeleteItem  func(uuid.UUID) error
	MockBillGetAllByUserID func(uuid.UUID) ([]model.Bill, error)
)

func NewBillStorageMock() storage.IBillStorage {
	return &BillStorageMock{}
}

type BillStorageMock struct {
}

func (b BillStorageMock) Create(bill model.Bill) (model.Bill, error) {
	return MockBillCreate(bill)
}

func (b BillStorageMock) UpdateBill(bill model.Bill) (model.Bill, error) {
	return MockBillUpdate(bill)
}

func (b BillStorageMock) GetByID(id uuid.UUID) (model.Bill, error) {
	return MockBillGetByID(id)
}

func (b BillStorageMock) DeleteBill(id uuid.UUID) error {
	return MockBillDelete(id)
}

func (b BillStorageMock) GetAllByUserID(userID uuid.UUID) ([]model.Bill, error) {
	return MockBillGetAllByUserID(userID)
}

func (b BillStorageMock) CreateItem(item model.Item) (model.Item, error) {
	return MockBillCreateItem(item)
}

func (b BillStorageMock) GetItemByID(id uuid.UUID) (model.Item, error) {
	return MockBillGetItemByID(id)
}

func (b BillStorageMock) UpdateItem(item model.Item) (model.Item, error) {
	return MockBillUpdateItem(item)
}

func (b BillStorageMock) DeleteItem(id uuid.UUID) error {
	return MockBillDeleteItem(id)
}
