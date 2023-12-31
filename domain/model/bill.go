package model

import (
	"github.com/google/uuid"
	"time"
)

type BillModel struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
	Date    time.Time
	Group   uuid.UUID
	Items   []ItemModel
}

func CreateBillModel(ownerID uuid.UUID, name string, date time.Time, group uuid.UUID, items []ItemModel) BillModel {
	return BillModel{
		ID:      uuid.New(),
		OwnerID: ownerID,
		Name:    name,
		Date:    date,
		Group:   group,
		Items:   items,
	}
}
