package dto

import (
	"errors"
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type GroupInputDTO struct {
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

func (g GroupInputDTO) ToGroup(owner model.User, members []model.User) model.Group {
	return model.CreateGroup(owner, g.Name, members)
}

func (g GroupInputDTO) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func ToGroupDTO(g *model.Group) GroupOutputDTO {
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
