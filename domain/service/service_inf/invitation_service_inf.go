package service_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IInvitationService interface {
	CreateGroupInvitation(request dto.GroupInputDTO, groupID uuid.UUID) error

	AcceptGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error

	DeclineGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error
}
