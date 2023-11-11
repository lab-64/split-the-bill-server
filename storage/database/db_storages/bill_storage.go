package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/storage_inf"
)

type BillStorage struct {
	DB *gorm.DB
}

func NewBillStorage(DB *database.Database) IBillStorage {
	return &BillStorage{DB: DB.Context}
}

func (b *BillStorage) Create(bill BillModel) error {
	//TODO implement me
	panic("implement me")
}

func (b *BillStorage) GetByID(id uuid.UUID) (BillModel, error) {
	//TODO implement me
	panic("implement me")
}
