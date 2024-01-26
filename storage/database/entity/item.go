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
	Contributors []*User   `gorm:"many2many:item_contributors;"` // many to many
}

// ToItemEntity converts an ItemModel to an Item
func ToItemEntity(item ItemModel) Item {
	// convert contributors to user entities
	var contributors []*User
	for _, contributor := range item.Contributors {
		contributors = append(contributors, &User{Base: Base{ID: contributor.ID}})
	}
	return Item{Base: Base{ID: item.ID}, Name: item.Name, Price: item.Price, BillID: item.BillID, Contributors: contributors}
}

// ToItemModel converts an Item to an ItemModel
func ToItemModel(item Item) ItemModel {
	// convert contributors to core user models
	var contributors []UserModel
	for _, contributor := range item.Contributors {
		contributors = append(contributors, ToCoreUserModel(*contributor))
	}
	return ItemModel{ID: item.ID, Name: item.Name, Price: item.Price, BillID: item.BillID, Contributors: contributors}
}
