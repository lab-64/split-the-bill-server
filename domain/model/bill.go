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
	Balance map[uuid.UUID]float64
}

func CreateBillModel(ownerID uuid.UUID, name string, date time.Time, groupID uuid.UUID, items []ItemModel) BillModel {
	return BillModel{
		ID:      uuid.New(),
		OwnerID: ownerID,
		Name:    name,
		Date:    date,
		GroupID: groupID,
		Items:   items,
	}
}

func (bill BillModel) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	for _, item := range bill.Items {
		ppp := item.Price / float64(len(item.Contributors))
		for _, contributor := range item.Contributors {
			balance[contributor] -= ppp
		}
		balance[bill.OwnerID] += item.Price
	}
	return balance
}
