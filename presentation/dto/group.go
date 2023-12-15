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
	Owner   UserCoreOutputDTO   `json:"owner"`
	ID      UUID                `json:"id"`
	Name    string              `json:"name"`
	Members []UserCoreOutputDTO `json:"members"`
	Bills   []BillCoreOutputDTO `json:"bills"`
}

func ToGroupModel(g GroupInputDTO) GroupModel {
	return CreateGroupModel(g.Owner, g.Name, []UUID{g.Owner})
}

func ToGroupDTO(g GroupModel) GroupOutputDTO {

	billsDTO := ToBillCoreDTOs(g.Bills)
	owner := ToUserCoreDTO(&g.Owner)
	members := ToUserCoreDTOs(g.Members)

	return GroupOutputDTO{
		Owner:   owner,
		ID:      g.ID,
		Name:    g.Name,
		Members: members,
		Bills:   billsDTO,
	}
}

// ValidateInput validates the inputs of the group creation request
func (g GroupInputDTO) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
