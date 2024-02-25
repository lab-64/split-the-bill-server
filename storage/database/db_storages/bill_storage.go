package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/converter"
	"split-the-bill-server/storage/database/entity"
)

type BillStorage struct {
	DB *gorm.DB
}

func NewBillStorage(DB *database.Database) storage.IBillStorage {
	return &BillStorage{DB: DB.Context}
}

func (b *BillStorage) Create(bill model.Bill) (model.Bill, error) {
	item := converter.ToBillEntity(bill)

	// store bill
	res := b.DB.Create(&item)
	if res.Error != nil {
		// TODO: does not trigger. Find out how to check for different errors
		// Adding a bill to an existing group but with a non-existing user as owner results in ErrForeignKeyViolated and therefore "NoSuchGroupError" is returned -> but in the fact user is missing
		if errors.Is(res.Error, gorm.ErrForeignKeyViolated) {
			return model.Bill{}, storage.NoSuchGroupError
		}
		// TODO: return general error
		return model.Bill{}, storage.NoSuchGroupError
	}
	// TODO: remove rows affected if we implement a fine error handling
	if res.RowsAffected == 0 {
		return model.Bill{}, storage.BillAlreadyExistsError
	}

	return converter.ToBillModel(item), res.Error
}

func (b *BillStorage) UpdateBill(bill model.Bill) (model.Bill, error) {
	billEntity := converter.ToBillEntity(bill)

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

	return converter.ToBillModel(billEntity), err
}

func (b *BillStorage) GetByID(id uuid.UUID) (model.Bill, error) {
	var bill entity.Bill
	tx := b.DB.Limit(1).Preload("Items.Contributors").Preload("Owner").Find(&bill, "id = ?", id)
	// TODO: return general error
	if tx.Error != nil {
		return model.Bill{}, storage.NoSuchBillError
	}
	if tx.RowsAffected == 0 {
		return model.Bill{}, storage.NoSuchBillError
	}
	billModel := converter.ToBillModel(bill)
	return billModel, nil
}

func (b *BillStorage) CreateItem(item model.Item) (model.Item, error) {
	itemEntity := converter.ToItemEntity(item)

	// TODO: if userId belongs to deleted user do not create
	// store item
	res := b.DB.
		Omit("Contributors.*"). // do not update user fields
		Preload("Contributors").
		Create(&itemEntity).
		First(&itemEntity, "id = ?", itemEntity.ID)

	// TODO: check if other errors can occur
	if res.Error != nil {
		return model.Item{}, storage.NoSuchBillError
	}
	if res.RowsAffected == 0 {
		return model.Item{}, storage.ItemAlreadyExistsError
	}

	return converter.ToItemModel(itemEntity), nil
}

func (b *BillStorage) GetItemByID(id uuid.UUID) (model.Item, error) {
	var item entity.Item
	tx := b.DB.Preload("Contributors").Limit(1).Find(&item, "id = ?", id)
	if tx.RowsAffected == 0 {
		return model.Item{}, storage.NoSuchItemError
	}
	return converter.ToItemModel(item), nil
}

// TODO: modify update method to only update the contributors list and not the whole item
func (b *BillStorage) UpdateItem(item model.Item) (model.Item, error) {
	itemEntity := converter.ToItemEntity(item)

	// run as a transaction to ensure consistency. item should be completely updated or not at all
	err := b.DB.Transaction(func(tx *gorm.DB) error {

		// update contributors associations
		res := tx.
			Omit("Contributors.*"). // do not update user fields
			Model(&itemEntity).
			Association("Contributors").
			Replace(itemEntity.Contributors)

		// TODO: add finer error handling
		if res != nil {
			return storage.NoSuchUserError
		}

		// update base item fields
		ret := tx.
			Omit(clause.Associations). // do not update associations
			Preload("Contributors").
			Model(&itemEntity).
			Updates(&itemEntity).
			First(&itemEntity)

		if ret.RowsAffected == 0 {
			return storage.NoSuchItemError
		}

		// TODO: add finer error handling
		if ret.Error != nil {
			return ret.Error
		}

		return nil
	})

	return converter.ToItemModel(itemEntity), err
}

func (b *BillStorage) DeleteItem(id uuid.UUID) error {
	res := b.DB.Delete(&entity.Item{}, "id = ?", id)
	if res.Error != nil {
		return storage.NoSuchBillError
	}
	return nil
}
