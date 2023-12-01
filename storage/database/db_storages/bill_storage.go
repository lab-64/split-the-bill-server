package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	. "split-the-bill-server/storage/storage_inf"
)

type BillStorage struct {
	DB *gorm.DB
}

func NewBillStorage(DB *database.Database) IBillStorage {
	return &BillStorage{DB: DB.Context}
}

func (b *BillStorage) Create(bill BillModel) error {
	item := ToBillEntity(bill)

	// store bill
	res := b.DB.Create(&item)
	if res.Error != nil {
		// TODO: does not trigger. Find out how to check for different errors
		if errors.Is(res.Error, gorm.ErrForeignKeyViolated) {
			return storage.NoSuchGroupError
		}
		// TODO: return general error
		return storage.NoSuchGroupError
	}
	// TODO: remove rows affected if we implement a fine error handling
	if res.RowsAffected == 0 {
		return storage.BillAlreadyExistsError
	}

	return res.Error
}

func (b *BillStorage) GetByID(id uuid.UUID) (BillModel, error) {
	var bill Bill
	tx := b.DB.Limit(1).Preload("Items").Find(&bill, "id = ?", id)
	// TODO: return general error
	if tx.Error != nil {
		return BillModel{}, storage.NoSuchBillError
	}
	if tx.RowsAffected == 0 {
		return BillModel{}, storage.NoSuchBillError
	}
	billModel := ToBillModel(bill)
	return billModel, nil
}

func (b *BillStorage) CreateItem(item ItemModel) (ItemModel, error) {
	itemEntity := ToItemEntity(item)

	// store item
	res := b.DB.Create(&itemEntity)
	if res.RowsAffected == 0 {
		return ItemModel{}, storage.BillAlreadyExistsError
	}

	return ToItemModel(itemEntity), res.Error
}
