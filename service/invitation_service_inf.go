package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/dto"
)

type IInvitationService interface {
	CreateGroupInvitation(request dto.GroupInvitationDTO) error

	AcceptGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error

	DeclineGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error
}
