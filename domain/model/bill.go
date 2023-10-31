package model

import (
	"github.com/google/uuid"
	"time"
)

type Bill struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
	Date    time.Time
	Items   []*Item
}

func CreateBill(ownerID uuid.UUID, name string, date time.Time, items []*Item) Bill {
	return Bill{
		ID:      uuid.New(),
		OwnerID: ownerID,
		Name:    name,
		Date:    date,
		Items:   items,
	}
}
