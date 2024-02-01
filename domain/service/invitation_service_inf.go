package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IInvitationService interface {
	CreateGroupInvitations(groupID uuid.UUID) (dto.GroupInvitationOutputDTO, error)

	GetGroupInvitationByID(invitationID uuid.UUID) (dto.GroupInvitationOutputDTO, error)

	AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error
}
