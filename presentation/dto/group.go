package dto

import (
	"errors"
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInputDTO struct {
	Owner   uuid.UUID   `json:"owner"`
	Name    string      `json:"name"`
	Invites []uuid.UUID `json:"invites"`
}

type GroupOutputDTO struct {
	Owner   uuid.UUID       `json:"owner"`
	ID      uuid.UUID       `json:"id"`
	Name    string          `json:"name"`
	Members []uuid.UUID     `json:"members"`
	Bills   []BillOutputDTO `json:"bills"`
}

func ToGroupModel(g GroupInputDTO) GroupModel {
	return CreateGroupModel(UserModel{ID: g.Owner}, g.Name)
}

func ToGroupDTO(g *GroupModel) GroupOutputDTO {
	billsDTO := make([]BillOutputDTO, len(g.Bills))

	for i, bill := range g.Bills {
		billsDTO[i] = ToBillDTO(bill)
	}
	// get all member ids
	var members []uuid.UUID
	for _, member := range g.Members {
		members = append(members, member.ID)
	}
	return GroupOutputDTO{
		Owner:   g.Owner.ID,
		ID:      g.ID,
		Name:    g.Name,
		Members: members,
		Bills:   billsDTO,
	}
}

// Validator

func (g GroupInputDTO) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
