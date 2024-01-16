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
