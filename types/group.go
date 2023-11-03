package types

import (
	"github.com/google/uuid"
)

type Group struct {
	Owner   uuid.UUID
	ID      uuid.UUID
	Name    string
	Members []uuid.UUID
	Bills   []*Bill
}

func CreateGroup(owner uuid.UUID, name string, members []uuid.UUID) Group {
	return Group{
		Owner:   owner,
		ID:      uuid.New(),
		Name:    name,
		Members: members,
		Bills:   make([]*Bill, 0),
	}
}
