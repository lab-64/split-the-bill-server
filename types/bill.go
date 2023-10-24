package types

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

type Item struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	Contributors []uuid.UUID
}

func CreateItem(name string, price float64, contributors []uuid.UUID) Item {
	return Item{
		ID:           uuid.New(),
		Name:         name,
		Price:        price,
		Contributors: contributors,
	}
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
