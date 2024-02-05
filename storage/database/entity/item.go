package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type Item struct {
	Base
	Name         string    `gorm:"not null"`
	Price        float64   `gorm:"not null"`
	BillID       uuid.UUID `gorm:"type:uuid"`                    // belongs to bill
	Contributors []*User   `gorm:"many2many:item_contributors;"` // many to many
}

func CreateItemEntity(item model.ItemModel) Item {
	// create user entities with the given ids for the contributors
	var contributors []*User
	for _, contributor := range item.Contributors {
		contributors = append(contributors, &User{Base: Base{ID: contributor.ID}})
	}
	return Item{Base: Base{ID: item.ID}, Name: item.Name, Price: item.Price, BillID: item.BillID, Contributors: contributors}
}

func ConvertToItemModel(item Item, isDetailed bool) model.ItemModel {
	contributors := make([]model.UserModel, len(item.Contributors))

	if isDetailed {
		for i, contributor := range item.Contributors {
			contributors[i] = ConvertToUserModel(*contributor, false)
		}
	}

	return model.ItemModel{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}
