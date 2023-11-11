package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IInvitationService interface {
	CreateGroupInvitation(request GroupInputDTO, groupID UUID) error

	AcceptGroupInvitation(invitation UUID, userID UUID) error

	DeclineGroupInvitation(invitation UUID, userID UUID) error
}
