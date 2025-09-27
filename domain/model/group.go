package model

import (
	"github.com/google/uuid"
	"sort"
	"split-the-bill-server/presentation/dto"
)

type Group struct {
	ID           uuid.UUID
	Name         string
	Owner        User
	Members      []User
	Bills        []Bill
	Balance      map[uuid.UUID]float64
	InvitationID uuid.UUID
}

func CreateGroup(id uuid.UUID, group dto.GroupInput, members []uuid.UUID) Group {

	// store memberIDs in empty User
	memberModel := make([]User, len(members))
	for i, member := range members {
		memberModel[i] = User{ID: member}
	}
	return Group{
		ID:           id,
		Owner:        User{ID: group.OwnerID}, // store ownerID in empty User
		Name:         group.Name,
		Members:      memberModel,
		InvitationID: uuid.New(),
	}
}

// CalculateBalance Calculate balance for each member in the group
func (group *Group) CalculateBalance() map[uuid.UUID]float64 {
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

// SortBillsByDate Sort bills by date descending
func (group *Group) SortBillsByDate() {
	sort.Slice(group.Bills, func(i, j int) bool {
		return group.Bills[i].Date.After(group.Bills[j].Date)
	})
}

func (group *Group) IsMember(memberID uuid.UUID) bool {
	isMember := false
	for _, member := range group.Members {
		if memberID == member.ID {
			isMember = true
			return isMember
		}
	}
	return isMember
}
