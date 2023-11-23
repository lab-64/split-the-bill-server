package model

import (
	"github.com/google/uuid"
	"time"
)

// TODO: how to handle item model if input dto does not include items but the output dto will include items?
type BillModel struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
	Date    time.Time
	Group   uuid.UUID
	Items   []ItemModel
}

func CreateBill(ownerID uuid.UUID, name string, date time.Time) BillModel {
	return BillModel{
		ID:      uuid.New(),
		OwnerID: ownerID,
		Name:    name,
		Date:    date,
	}
}
