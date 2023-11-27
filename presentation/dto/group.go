package dto

import (
	"errors"
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInputDTO struct {
	Owner UUID   `json:"ownerID"`
	Name  string `json:"name"`
}

type GroupOutputDTO struct {
	Owner   UUID   `json:"ownerID"`
	ID      UUID   `json:"id"`
	Name    string `json:"name"`
	Members []UUID `json:"memberIDs"`
	Bills   []UUID `json:"billIDs"`
}

func ToGroupModel(g GroupInputDTO, members []UUID) GroupModel {
	return CreateGroupModel(g.Owner, g.Name, members)
}

func ToGroupDTO(g GroupModel) GroupOutputDTO {

	return GroupOutputDTO{
		Owner:   g.Owner,
		ID:      g.ID,
		Name:    g.Name,
		Members: g.Members,
		Bills:   g.Bills,
	}
}

// ValidateInput validates the inputs of the group creation request
func (g GroupInputDTO) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
