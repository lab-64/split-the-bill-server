package dto

import (
	"errors"
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInputDTO struct {
	OwnerID UUID   `json:"ownerID"`
	Name    string `json:"name"`
}

type GroupCoreOutputDTO struct {
	Owner   UserCoreOutputDTO   `json:"owner"`
	ID      UUID                `json:"id"`
	Name    string              `json:"name"`
	Members []UserCoreOutputDTO `json:"members"`
}

type GroupDetailedOutputDTO struct {
	Owner   UserCoreOutputDTO       `json:"owner"`
	ID      UUID                    `json:"id"`
	Name    string                  `json:"name"`
	Members []UserCoreOutputDTO     `json:"members"`
	Bills   []BillDetailedOutputDTO `json:"bills"`
}

func ToGroupModel(g GroupInputDTO) GroupModel {
	return CreateGroupModel(g.OwnerID, g.Name, []UUID{g.OwnerID})
}

func ToGroupCoreDTO(g GroupModel) GroupCoreOutputDTO {

	owner := ToUserCoreDTO(&g.Owner)
	members := ToUserCoreDTOs(g.Members)

	return GroupCoreOutputDTO{
		Owner:   owner,
		ID:      g.ID,
		Name:    g.Name,
		Members: members,
	}
}

func ToGroupDetailedDTO(g GroupModel) GroupDetailedOutputDTO {

	billsDTO := ToBillDetailedDTOs(g.Bills)
	owner := ToUserCoreDTO(&g.Owner)
	members := ToUserCoreDTOs(g.Members)

	return GroupDetailedOutputDTO{
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
