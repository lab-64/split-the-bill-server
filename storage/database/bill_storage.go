package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type BillStorage struct {
	DB *gorm.DB
}

func NewBillStorage(DB *Database) storage.IBillStorage {
	return &BillStorage{DB: DB.context}
}

func (b *BillStorage) Create(bill types.Bill) error {
	//TODO implement me
	panic("implement me")
}

func (b *BillStorage) GetByID(id uuid.UUID) (types.Bill, error) {
	//TODO implement me
	panic("implement me")
}
