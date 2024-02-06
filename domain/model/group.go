package model

import (
	"github.com/google/uuid"
)

type GroupModel struct {
	ID           uuid.UUID
	Name         string
	Owner        UserModel
	Members      []UserModel
	Bills        []BillModel
	Balance      map[uuid.UUID]float64
	InvitationID uuid.UUID
}

func (group *GroupModel) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	// init balance for all members
	for _, member := range group.Members {
		balance[member.ID] = 0
	}
	for i, bill := range group.Bills {
		billBalance := bill.CalculateBalance()
		// update group balance
		for k, v := range billBalance {
			balance[k] += v
		}
		// set balance for each bill
		group.Bills[i].Balance = billBalance
	}
	return balance
}
