package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IInvitationService interface {
	CreateGroupInvitation(request GroupInvitationInputDTO) error

	AcceptGroupInvitation(invitation UUID, userID UUID) error

	DeclineGroupInvitation(invitation UUID, userID UUID) error

	GetGroupInvitationByID(id UUID) (GroupInvitationOutputDTO, error)

	GetGroupInvitationsFromUser(id UUID) ([]GroupInvitationOutputDTO, error)
}
