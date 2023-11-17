package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Bill struct {
	Base
	Name    string    `gorm:"not null"`
	Items   []Item    `gorm:"foreignKey:BillID"` // has many items
	GroupID uuid.UUID `gorm:"type:uuid"`         // belongs to group
}

func ToBillEntity(bill BillModel) Bill {
	// convert items
	var items []Item
	for _, item := range bill.Items {
		items = append(items, ToItemEntity(item))
	}
	return Bill{Base: Base{ID: bill.ID}, Name: bill.Name, Items: items, GroupID: bill.Group}
}
