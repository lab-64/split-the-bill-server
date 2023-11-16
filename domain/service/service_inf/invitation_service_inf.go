package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IInvitationService interface {
	CreateGroupInvitation(request GroupInvitationInputDTO) error

	AcceptGroupInvitation(invitationID UUID, userID UUID) error

	DeclineGroupInvitation(invitationID UUID) error

	GetGroupInvitationByID(invitationID UUID) (GroupInvitationOutputDTO, error)

	GetGroupInvitationsByUser(userID UUID) ([]GroupInvitationOutputDTO, error)
}
