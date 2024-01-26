package model

import (
	"github.com/google/uuid"
	"time"
)

type BillModel struct {
	ID      uuid.UUID
	Owner   UserModel
	Name    string
	Date    time.Time
	GroupID uuid.UUID
	Items   []ItemModel
}

func CreateBillModel(owner UserModel, name string, date time.Time, groupID uuid.UUID, items []ItemModel) BillModel {
	return BillModel{
		ID:      uuid.New(),
		Owner:   owner,
		Name:    name,
		Date:    date,
		GroupID: groupID,
		Items:   items,
	}
}
