package dto

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type GroupInvitationInputDTO struct {
	Issuer  UUID   `json:"issuer"`
	GroupID UUID   `json:"groupID"`
	Invites []UUID `json:"invites"`
}

type GroupInvitationOutputDTO struct {
	InvitationID UUID           `json:"invitationID"`
	Group        GroupOutputDTO `json:"group"`
}

type HandleInvitationInputDTO struct {
	Issuer       UUID `json:"issuer"`
	InvitationID UUID `json:"invitationID"`
}

func ToGroupInvitationDTO(invitation model.GroupInvitationModel) GroupInvitationOutputDTO {
	group := ToGroupDTO(invitation.Group)
	return GroupInvitationOutputDTO{InvitationID: invitation.ID, Group: group}
}
