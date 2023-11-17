package model

import (
	"github.com/google/uuid"
)

type GroupModel struct {
	ID      uuid.UUID
	Owner   uuid.UUID
	Name    string
	Members []uuid.UUID
	// TODO: Decide which struct (group or bill) should have a full reference to the other
	Bills []BillModel
}

func CreateGroupModel(owner uuid.UUID, name string, members []uuid.UUID) GroupModel {
	return GroupModel{
		Owner:   owner,
		ID:      uuid.New(),
		Name:    name,
		Members: members,
		Bills:   make([]BillModel, 0),
	}
}
