package dto

import (
	"errors"
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type GroupCreateDTO struct {
	Name    string      `json:"name"`
	Invites []uuid.UUID `json:"invites"`
}

type GroupDTO struct {
	Owner   uuid.UUID   `json:"owner"`
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	Members []uuid.UUID `json:"members"`
	Bills   []BillDTO   `json:"bills"`
}

func (g GroupCreateDTO) ToGroup(owner uuid.UUID, members []uuid.UUID) types.Group {
	return types.CreateGroup(owner, g.Name, members)
}

func (g GroupCreateDTO) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func ToGroupDTO(g *types.Group) GroupDTO {
	billsDTO := make([]BillDTO, len(g.Bills))

	for i, bill := range g.Bills {
		billsDTO[i] = ToBillDTO(bill)
	}

	return GroupDTO{
		Owner:   g.Owner,
		ID:      g.ID,
		Name:    g.Name,
		Members: g.Members,
		Bills:   billsDTO,
	}
}
