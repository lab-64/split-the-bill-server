package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IInvitationService interface {
	CreateGroupInvitations(request GroupInvitationInputDTO) ([]GroupInvitationOutputDTO, error)

	HandleGroupInvitation(invitationID UUID, isAccept bool) error

	GetGroupInvitationByID(invitationID UUID) (GroupInvitationOutputDTO, error)

	GetGroupInvitationsByUser(userID UUID) ([]GroupInvitationOutputDTO, error)
}
