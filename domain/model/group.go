package model

import (
	"github.com/google/uuid"
)

type GroupModel struct {
	ID      uuid.UUID
	Name    string
	Owner   UserModel
	Members []UserModel
	Bills   []BillModel
	Balance map[uuid.UUID]float64
}

func (group GroupModel) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	// init balance for all members
	for _, member := range group.Members {
		balance[member.ID] = 0
	}
	for _, bill := range group.Bills {
		for _, item := range bill.Items {
			ppp := item.Price / float64(len(item.Contributors))
			for _, contributor := range item.Contributors {
				balance[contributor.ID] -= ppp
			}
			balance[bill.Owner.ID] += item.Price
		}
	}
	return balance
}
