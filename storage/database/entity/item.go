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

func ToItemEntity(item ItemModel) Item {
	return Item{Base: Base{ID: item.ID}, Name: item.Name, Price: item.Price}
}

func ToItemModel(item *Item) ItemModel {
	return ItemModel{ID: item.ID, Name: item.Name, Price: item.Price}
}
