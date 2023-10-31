package model

import (
	"github.com/google/uuid"
)

type Group struct {
	Owner   User
	ID      uuid.UUID
	Name    string
	Members []User
	Bills   []*Bill
}

func CreateGroup(owner User, name string, members []User) Group {
	return Group{
		Owner:   owner,
		ID:      uuid.New(),
		Name:    name,
		Members: members,
		Bills:   make([]*Bill, 0),
	}
}
