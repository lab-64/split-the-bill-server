package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInvitationInputDTO struct {
	Issuer   uuid.UUID   `json:"issuerID"`
	GroupID  uuid.UUID   `json:"groupID"`
	Invitees []uuid.UUID `json:"inviteeIDs"`
}

type GroupInvitationOutputDTO struct {
	InvitationID uuid.UUID          `json:"invitationID"`
	Group        GroupCoreOutputDTO `json:"group"`
}

type InvitationResponseInputDTO struct {
	IsAccept bool `json:"isAccept"`
}

func CreateGroupInvitationModel(id uuid.UUID, groupID uuid.UUID) GroupInvitationModel {
	return GroupInvitationModel{
		ID:    id,
		Group: GroupModel{ID: groupID},
	}
}

func ConvertToGroupInvitationDTO(invitation GroupInvitationModel) GroupInvitationOutputDTO {
	group := ConvertToGroupCoreDTO(invitation.Group)

	return GroupInvitationOutputDTO{
		InvitationID: invitation.ID,
		Group:        group,
	}
}
