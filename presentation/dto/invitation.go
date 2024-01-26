package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
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

func CreateGroupInvitationModel(id uuid.UUID, groupID uuid.UUID, inviteeID uuid.UUID) GroupInvitationModel {
	return GroupInvitationModel{
		ID:      id,
		Date:    time.Now(),
		Group:   GroupModel{ID: groupID},
		Invitee: UserModel{ID: inviteeID},
	}
}

func ConvertToGroupInvitationDTO(invitation GroupInvitationModel) GroupInvitationOutputDTO {
	group := ConvertToGroupCoreDTO(invitation.Group)

	return GroupInvitationOutputDTO{
		InvitationID: invitation.ID,
		Group:        group,
	}
}
