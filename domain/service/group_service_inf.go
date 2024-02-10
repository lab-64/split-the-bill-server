package service

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IGroupService interface {
	Create(groupDTO GroupInputDTO) (GroupDetailedOutputDTO, error)

	Update(userID UUID, groupID UUID, group GroupInputDTO) (GroupDetailedOutputDTO, error)

	GetByID(id UUID) (GroupDetailedOutputDTO, error)

	GetAll(userID UUID, invitationID UUID) ([]GroupDetailedOutputDTO, error)

	AcceptGroupInvitation(invitationID UUID, userID UUID) error
}
