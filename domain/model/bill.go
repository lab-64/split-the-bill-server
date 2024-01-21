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
	GroupID uuid.UUID
	Items   []ItemModel
}

func CreateBillModel(id uuid.UUID, owner uuid.UUID, name string, data time.Time, groupID uuid.UUID, items []ItemModel) BillModel {
	return BillModel{
		ID:      id,
		OwnerID: owner,
		Name:    name,
		Date:    data,
		GroupID: groupID,
		Items:   items,
	}
}
