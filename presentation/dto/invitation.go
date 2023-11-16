package dto

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type GroupInvitationInputDTO struct {
	Issuer   UUID   `json:"issuer"`
	GroupID  UUID   `json:"groupID"`
	Invitees []UUID `json:"invitees"`
}

type GroupInvitationOutputDTO struct {
	InvitationID UUID           `json:"invitationID"`
	Group        GroupOutputDTO `json:"group"`
}

type HandleInvitationInputDTO struct {
	Issuer       UUID `json:"issuer"`
	InvitationID UUID `json:"invitationID"`
}

func ToGroupInvitationModel(groupID UUID, userID UUID) model.GroupInvitationModel {
	return model.CreateGroupInvitation(groupID, userID)
}

func ToGroupInvitationDTO(invitation model.GroupInvitationModel) GroupInvitationOutputDTO {
	group := ToGroupDTO(invitation.Group)
	return GroupInvitationOutputDTO{InvitationID: invitation.ID, Group: group}
}
