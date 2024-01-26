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

func CreateGroupModel(owner uuid.UUID, name string, members []uuid.UUID) GroupModel {
	// store memberIDs in empty UserModel
	memberModel := make([]UserModel, len(members))
	for i, member := range members {
		memberModel[i] = UserModel{ID: member}
	}
	return GroupModel{
		ID:      uuid.New(),
		Owner:   UserModel{ID: owner}, // store ownerID in empty UserModel
		Name:    name,
		Members: memberModel,
	}
}

func (group GroupModel) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	// init balance for all members
	for _, member := range group.Members {
		balance[member.ID] = 0
	}
	for _, bill := range group.Bills {
		for k, v := range bill.CalculateBalance() {
			balance[k] += v
		}
	}
	return balance
}
