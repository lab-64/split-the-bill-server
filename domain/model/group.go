package model

import (
	"github.com/google/uuid"
)

type GroupModel struct {
	Owner   UserModel
	ID      uuid.UUID
	Name    string
	Members []UserModel
	Bills   []BillModel
}

func CreateGroupModel(owner UserModel, name string) GroupModel {
	return GroupModel{
		Owner:   owner,
		ID:      uuid.New(),
		Name:    name,
		Members: make([]UserModel, 0),
		Bills:   make([]BillModel, 0),
	}
}
