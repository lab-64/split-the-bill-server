package dto

import (
	"errors"
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInputDTO struct {
	OwnerID uuid.UUID `json:"ownerID"`
	Name    string    `json:"name"`
}

type GroupCoreOutputDTO struct {
	Owner   UserCoreOutputDTO   `json:"owner"`
	ID      uuid.UUID           `json:"id"`
	Name    string              `json:"name"`
	Members []UserCoreOutputDTO `json:"members"`
}

type GroupDetailedOutputDTO struct {
	Owner        UserCoreOutputDTO       `json:"owner"`
	ID           uuid.UUID               `json:"id"`
	Name         string                  `json:"name"`
	Members      []UserCoreOutputDTO     `json:"members"`
	Bills        []BillDetailedOutputDTO `json:"bills"`
	Balance      map[uuid.UUID]float64   `json:"balance,omitempty"`      // include balance only if balance is set
	InvitationID uuid.UUID               `json:"invitationID,omitempty"` // include invitationID only if invitationID is set
}

func CreateGroupModel(id uuid.UUID, group GroupInputDTO, members []uuid.UUID) GroupModel {

	// store memberIDs in empty UserModel
	memberModel := make([]UserModel, len(members))
	for i, member := range members {
		memberModel[i] = UserModel{ID: member}
	}
	return GroupModel{
		ID:           id,
		Owner:        UserModel{ID: group.OwnerID}, // store ownerID in empty UserModel
		Name:         group.Name,
		Members:      memberModel,
		InvitationID: uuid.New(),
	}
}

func ConvertToGroupDetailedDTO(g GroupModel) GroupDetailedOutputDTO {

	billsDTO := ConvertToBillDetailedDTOs(g.Bills)
	owner := ConvertToUserCoreDTO(&g.Owner)
	members := ConvertToUserCoreDTOs(g.Members)
	return GroupDetailedOutputDTO{
		Owner:        owner,
		ID:           g.ID,
		Name:         g.Name,
		Members:      members,
		Bills:        billsDTO,
		Balance:      g.Balance,
		InvitationID: g.InvitationID,
	}
}

// ValidateInput validates the inputs of the group creation request
func (g GroupInputDTO) ValidateInput() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
