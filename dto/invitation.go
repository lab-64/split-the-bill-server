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

type InvitationInputDTO struct {
	Type   string    `json:"type"`
	ID     uuid.UUID `json:"id"`
	Accept bool      `json:"accept"`
}

func ToGroupInvitationDTO(invitation types.GroupInvitation) GroupInvitationOutputDTO {
	group := ToGroupDTO(invitation.Group)
	return GroupInvitationOutputDTO{InvitationID: invitation.ID, Group: group}
}
