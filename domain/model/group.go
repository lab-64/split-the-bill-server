package model

import (
	"github.com/google/uuid"
)

type GroupModel struct {
	ID      uuid.UUID
	Owner   uuid.UUID
	Name    string
	Members []uuid.UUID
	Bills   []BillModel
}

func CreateGroupModel(owner uuid.UUID, name string, members []uuid.UUID) GroupModel {
	return GroupModel{
		Owner:   owner,
		ID:      uuid.New(),
		Name:    name,
		Members: members,
	}
}
