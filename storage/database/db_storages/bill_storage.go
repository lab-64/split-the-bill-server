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
			Omit(clause.Associations).
			Preload("Owner").
			Model(&billEntity).
			Updates(&billEntity).
			First(&billEntity)

		if ret.RowsAffected == 0 {
			return storage.NoSuchBillError
		}

		if ret.Error != nil {
			return ret.Error
		}

		// update unseenFrom associations
		err := tx.Model(&billEntity).
			Omit(clause.Associations).
			Association("UnseenFrom").
			Replace(billEntity.UnseenFrom)

		// delete old items
		err = tx.
			Where("bill_id = ?", billEntity.ID).
			Delete(&entity.Item{}).
			Error

		// insert new items with contributors
		for _, item := range billEntity.Items {
			err = tx.
				Preload("Contributors").
				Create(&item).
				First(&item, "id = ?", item.ID).
				Error
		}

		return err
	})

	return converter.ToBillModel(billEntity), err
}

func (b *BillStorage) GetByID(id uuid.UUID) (model.Bill, error) {
	var bill entity.Bill
	tx := b.DB.Limit(1).
		Preload("Items.Contributors").
		Preload("Owner").
		Preload("UnseenFrom").
		Find(&bill, "id = ?", id)
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

func (b *BillStorage) DeleteBill(id uuid.UUID) error {
	bill := entity.Bill{
		Base: entity.Base{ID: id},
	}

	res := b.DB.Select(clause.Associations).Delete(&bill)
	if res.Error != nil {
		return storage.NoSuchBillError
	}
	return nil
}
