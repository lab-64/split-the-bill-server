package db_storages

import (
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
	if res.RowsAffected == 0 {
		return storage.BillAlreadyExistsError
	}

	return res.Error
}

func (b *BillStorage) GetByID(id uuid.UUID) (BillModel, error) {
	var bill Bill
	tx := b.DB.Limit(1).Preload("Items.Contributors").Find(&bill, "id = ?", id)
	if tx.RowsAffected == 0 {
		return BillModel{}, storage.NoSuchBillError
	}
	// TODO: return general error
	if tx.Error != nil {
		return BillModel{}, storage.NoSuchBillError
	}
	billModel := ToBillModel(bill)
	return billModel, nil
}

func (b *BillStorage) CreateItem(item ItemModel) (ItemModel, error) {
	itemEntity := ToItemEntity(item)

	// store item
	res := b.DB.Create(&itemEntity)

	// TODO: check if other errors can occur
	if res.Error != nil {
		return ItemModel{}, storage.NoSuchUserError
	}
	if res.RowsAffected == 0 {
		return ItemModel{}, storage.ItemAlreadyExistsError
	}

	return ToItemModel(itemEntity), nil
}

func (b *BillStorage) GetItemByID(id uuid.UUID) (ItemModel, error) {
	var item Item
	tx := b.DB.Preload("Contributors").Limit(1).Find(&item, "id = ?", id)
	if tx.RowsAffected == 0 {
		return ItemModel{}, storage.NoSuchItemError
	}
	return ToItemModel(item), nil
}

func (b *BillStorage) UpdateItem(item ItemModel) (ItemModel, error) {
	itemEntity := ToItemEntity(item)

	// store item
	res := b.DB.Model(&itemEntity).Updates(&itemEntity)

	// TODO: check if other errors can occur
	if res.Error != nil {
		return ItemModel{}, storage.NoSuchUserError
	}

	// TODO: return the stored model instead of the input model
	return ToItemModel(itemEntity), nil
}
