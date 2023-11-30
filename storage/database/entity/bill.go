package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type Bill struct {
	Base
	OwnerUID uuid.UUID `gorm:"type:uuid"`
	Owner    User      `gorm:"foreignKey:OwnerUID"` // belongs to user
	Name     string    `gorm:"not null"`
	Date     time.Time
	Items    []Item    `gorm:"foreignKey:BillID"` // has many items
	GroupID  uuid.UUID `gorm:"type:uuid"`         // group has many bills
}

// ToBillEntity converts a BillModel to a Bill
func ToBillEntity(bill BillModel) Bill {

	// convert items
	var items []Item
	for _, item := range bill.Items {
		items = append(items, ToItemEntity(item))
	}

	return Bill{
		Base:     Base{ID: bill.ID},
		Name:     bill.Name,
		Date:     bill.Date,
		GroupID:  bill.Group,
		OwnerUID: bill.OwnerID,
		Items:    items}
}

// ToBillModel converts a Bill to a BillModel
func ToBillModel(bill Bill) BillModel {

	// convert items
	var items []ItemModel
	for _, item := range bill.Items {
		items = append(items, ToItemModel(item))
	}

	return BillModel{ID: bill.ID, Name: bill.Name, Date: bill.Date, Group: bill.GroupID, Items: items, OwnerID: bill.OwnerUID}
}
