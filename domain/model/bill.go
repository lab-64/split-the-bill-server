package model

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
	"time"
)

type Bill struct {
	ID      uuid.UUID
	Owner   User
	Name    string
	Date    time.Time
	GroupID uuid.UUID
	Items   []Item
	Balance map[uuid.UUID]float64
}

func CreateBill(id uuid.UUID, b dto.BillInput, date time.Time) Bill {
	// convert each item
	var items []Item
	for _, item := range b.Items {
		items = append(items, CreateItem(uuid.New(), item))
	}
	return Bill{
		ID:      id,
		Owner:   User{ID: b.OwnerID},
		Name:    b.Name,
		Date:    date,
		GroupID: b.GroupID,
		Items:   items,
	}
}

func (bill *Bill) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	for _, item := range bill.Items {
		ppp := item.Price / float64(len(item.Contributors))
		for _, contributor := range item.Contributors {
			balance[contributor.ID] -= ppp
		}
		balance[bill.Owner.ID] += item.Price
	}
	return balance
}
