package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Item struct {
	Base
	Name   string    `gorm:"not null"`
	Price  float64   `gorm:"not null"`
	BillID uuid.UUID `gorm:"type:uuid"` // belongs to bill
}

// ToItemEntity converts an ItemModel to an Item
func ToItemEntity(item ItemModel) Item {
	return Item{Base: Base{ID: item.ID}, Name: item.Name, Price: item.Price, BillID: item.BillID}
}

// ToItemModel converts an Item to an ItemModel
func ToItemModel(item Item) ItemModel {
	return ItemModel{ID: item.ID, Name: item.Name, Price: item.Price, BillID: item.BillID}
}
