package dto

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInvitationInputDTO struct {
	Issuer   UUID   `json:"issuerID"`
	GroupID  UUID   `json:"groupID"`
	Invitees []UUID `json:"inviteeIDs"`
}

type GroupInvitationOutputDTO struct {
	InvitationID UUID               `json:"invitationID"`
	Group        GroupCoreOutputDTO `json:"group"`
}

type InvitationResponseInputDTO struct {
	IsAccept bool `json:"isAccept"`
}

func ToGroupInvitationModel(groupID UUID, userID UUID) GroupInvitationModel {
	return CreateGroupInvitation(groupID, userID)
}

func ToGroupInvitationDTO(invitation GroupInvitationModel) GroupInvitationOutputDTO {
	group := ToGroupCoreDTO(invitation.Group)

	return GroupInvitationOutputDTO{
		InvitationID: invitation.ID,
		Group:        group,
	}
}
