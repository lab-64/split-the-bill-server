package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
)

type BillStorage struct {
	DB *gorm.DB
}

func NewBillStorage(DB *database.Database) storage.IBillStorage {
	return &BillStorage{DB: DB.Context}
}

func (b *BillStorage) Create(bill BillModel) (BillModel, error) {
	item := CreateBillEntity(bill)

	// store bill
	res := b.DB.Create(&item)
	if res.Error != nil {
		// TODO: does not trigger. Find out how to check for different errors
		// Adding a bill to an existing group but with a non-existing user as owner results in ErrForeignKeyViolated and therefore "NoSuchGroupError" is returned -> but in the fact user is missing
		if errors.Is(res.Error, gorm.ErrForeignKeyViolated) {
			return BillModel{}, storage.NoSuchGroupError
		}
		// TODO: return general error
		return BillModel{}, storage.NoSuchGroupError
	}
	// TODO: remove rows affected if we implement a fine error handling
	if res.RowsAffected == 0 {
		return BillModel{}, storage.BillAlreadyExistsError
	}

	return ConvertToBillModel(item), res.Error
}

func (b *BillStorage) UpdateBill(bill BillModel) (BillModel, error) {
	billEntity := CreateBillEntity(bill)

	err := b.DB.Transaction(func(tx *gorm.DB) error {
		// update base bill fields
		ret := b.DB.
			Model(&billEntity).
			Updates(&billEntity)

		if ret.RowsAffected == 0 {
			return storage.NoSuchBillError
		}

		if ret.Error != nil {
			return ret.Error
		}

		// update items
		res := tx.
			Model(&billEntity).
			Association("Items").
			Replace(billEntity.Items)

		if res != nil {
			return res
		}

		return nil
	})

	return ConvertToBillModel(billEntity), err
}

func (b *BillStorage) GetByID(id uuid.UUID) (BillModel, error) {
	var bill Bill
	tx := b.DB.Limit(1).Preload("Items.Contributors").Preload("Owner").Find(&bill, "id = ?", id)
	// TODO: return general error
	if tx.Error != nil {
		return BillModel{}, storage.NoSuchBillError
	}
	if tx.RowsAffected == 0 {
		return BillModel{}, storage.NoSuchBillError
	}
	billModel := ConvertToBillModel(bill)
	return billModel, nil
}

func (b *BillStorage) GetAllByUserID(userID uuid.UUID) ([]BillModel, error) {
	var bills []Bill

	tx := b.DB.
		Preload(clause.Associations).
		Preload("Items.Contributors").
		Preload("Owner").
		Where("owner_id = ?", userID).
		Find(&bills)

	if tx.Error != nil {
		return nil, storage.NoSuchBillError
	}

	return ToBillModelSlice(bills), nil
}

func (b *BillStorage) CreateItem(item ItemModel) (ItemModel, error) {
	itemEntity := CreateItemEntity(item)

	// TODO: if userId belongs to deleted user do not create
	// store item
	res := b.DB.Create(&itemEntity)

	// TODO: check if other errors can occur
	if res.Error != nil {
		return ItemModel{}, storage.NoSuchUserError
	}
	if res.RowsAffected == 0 {
		return ItemModel{}, storage.ItemAlreadyExistsError
	}

	return ConvertToItemModel(itemEntity), nil
}

func (b *BillStorage) GetItemByID(id uuid.UUID) (ItemModel, error) {
	var item Item
	tx := b.DB.Preload("Contributors").Limit(1).Find(&item, "id = ?", id)
	if tx.RowsAffected == 0 {
		return ItemModel{}, storage.NoSuchItemError
	}
	return ConvertToItemModel(item), nil
}

func (b *BillStorage) UpdateItem(item ItemModel) (ItemModel, error) {
	itemEntity := CreateItemEntity(item)

	// run as a transaction to ensure consistency. item should be completely updated or not at all
	err := b.DB.Transaction(func(tx *gorm.DB) error {
		// update base item fields
		ret := tx.
			Model(&itemEntity).
			Updates(&itemEntity)

		if ret.RowsAffected == 0 {
			return storage.NoSuchItemError
		}

		// TODO: add finer error handling
		if ret.Error != nil {
			return ret.Error
		}

		// update contributors associations
		res := tx.
			Model(&itemEntity).
			Association("Contributors").
			Replace(itemEntity.Contributors)

		// TODO: add finer error handling
		if res != nil {
			return storage.NoSuchUserError
		}

		return nil
	})

	return ConvertToItemModel(itemEntity), err
}
