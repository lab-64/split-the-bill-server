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
	Balance map[uuid.UUID]float64
}

func (bill *BillModel) CalculateBalance() map[uuid.UUID]float64 {
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
