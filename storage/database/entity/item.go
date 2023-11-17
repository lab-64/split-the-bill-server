package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Item struct {
	Base
	Name         string    `gorm:"not null"`
	Price        float64   `gorm:"not null"`
	BillID       uuid.UUID `gorm:"type:uuid"`                   // belongs to bill
	Contributors []*User   `gorm:"many2many:item_contributors"` // many to many
}

func ToItemEntity(item ItemModel) Item {
	// convert uuids to users
	var contributors []*User
	for _, contributor := range item.Contributors {
		contributors = append(contributors, &User{Base: Base{ID: contributor}})
	}
	return Item{Base: Base{ID: item.ID}, Name: item.Name, Price: item.Price, Contributors: contributors}
}

func ToItemModel(item *Item) ItemModel {
	// convert users to uuids
	var contributors []uuid.UUID
	for _, contributor := range item.Contributors {
		contributors = append(contributors, contributor.ID)
	}
	return ItemModel{ID: item.ID, Name: item.Name, Price: item.Price, Contributors: contributors}
}
