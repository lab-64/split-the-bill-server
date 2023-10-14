package wire

import (
	"errors"
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type Group struct {
	Name    string      `json:"name"`
	Invites []uuid.UUID `json:"invites"`
}

func (g Group) ToGroup(owner uuid.UUID, members []uuid.UUID) types.Group {
	return types.CreateGroup(owner, g.Name, members)
}

// TODO: maybe move to types/group.go
func (g Group) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
