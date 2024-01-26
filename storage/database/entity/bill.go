package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"time"
)

type Bill struct {
	Base
	OwnerID uuid.UUID `gorm:"type:uuid"`
	Owner   User      `gorm:"foreignKey:OwnerID"` // belongs to user
	Name    string    `gorm:"not null"`
	Date    time.Time
	Items   []Item    `gorm:"foreignKey:BillID"` // has many items
	GroupID uuid.UUID `gorm:"type:uuid"`         // group has many bills
}

func CreateBillEntity(bill model.BillModel) Bill {

	// convert items
	var items []Item
	for _, item := range bill.Items {
		items = append(items, CreateItemEntity(item))
	}

	return Bill{
		Base:    Base{ID: bill.ID},
		Name:    bill.Name,
		Date:    bill.Date,
		GroupID: bill.GroupID,
		OwnerID: bill.Owner.ID,
		Items:   items,
	}
}

func ConvertToBillModel(bill Bill, isDetailed bool) model.BillModel {
	items := make([]model.ItemModel, len(bill.Items))
	owner := model.UserModel{}

	if isDetailed {

		for i, item := range bill.Items {
			items[i] = ConvertToItemModel(item, true)
		}
		owner = ConvertToUserModel(bill.Owner, false)
	}

	return model.BillModel{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Owner:   owner,
		GroupID: bill.GroupID,
		Items:   items,
	}
}
