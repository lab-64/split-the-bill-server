package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Item struct {
	Base
	Name         string    `gorm:"not null"`
	Price        float64   `gorm:"not null"`
	BillID       uuid.UUID `gorm:"type:uuid"`                    // belongs to bill
	Contributors []User    `gorm:"many2many:item_contributors;"` // many to many
}

// ToItemEntity converts an ItemModel to an Item
func ToItemEntity(item ItemModel) Item {
	// convert contributor uuids to users
	var contributors []User
	for _, contributor := range item.Contributors {
		contributors = append(contributors, User{Base: Base{ID: contributor}})
	}
	return Item{Base: Base{ID: item.ID}, Name: item.Name, Price: item.Price, BillID: item.BillID, Contributors: contributors}
}

// ToItemModel converts an Item to an ItemModel
func ToItemModel(item Item) ItemModel {
	// convert contributors to uuids
	var contributors []uuid.UUID
	for _, contributor := range item.Contributors {
		contributors = append(contributors, contributor.ID)
	}
	return ItemModel{ID: item.ID, Name: item.Name, Price: item.Price, BillID: item.BillID, Contributors: contributors}
}
