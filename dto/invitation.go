package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type GroupInvitationDTO struct {
	Issuer   uuid.UUID   `json:"issuer"`
	GroupID  uuid.UUID   `json:"groupID"`
	Invitees []uuid.UUID `json:"invitees"`
}

type GroupInvitationOutputDTO struct {
	InvitationID uuid.UUID      `json:"invitationID"`
	Group        GroupOutputDTO `json:"group"`
}

type HandleInvitationInputDTO struct {
	Issuer       uuid.UUID `json:"issuer"`
	InvitationID uuid.UUID `json:"invitationID"`
}

func ToGroupInvitationDTO(invitation types.GroupInvitation) GroupInvitationOutputDTO {
	group := ToGroupDTO(invitation.Group)
	return GroupInvitationOutputDTO{InvitationID: invitation.ID, Group: group}
}
