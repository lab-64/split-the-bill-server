package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type Bill struct {
	Base
	Name    string `gorm:"not null"`
	Date    time.Time
	Items   []Item    `gorm:"foreignKey:BillID"` // has many items
	GroupID uuid.UUID `gorm:"type:uuid"`         // belongs to group
}

// ToBillEntity converts a BillModel to a Bill
func ToBillEntity(bill BillModel) Bill {
	return Bill{Base: Base{ID: bill.ID}, Name: bill.Name, Date: bill.Date, GroupID: bill.Group}
}

// ToBillModel converts a Bill to a BillModel
func ToBillModel(bill Bill) BillModel {

	// convert items
	var items []ItemModel
	for _, item := range bill.Items {
		items = append(items, ToItemModel(item))
	}

	return BillModel{ID: bill.ID, Name: bill.Name, Date: bill.Date, Group: bill.GroupID, Items: items}
}
