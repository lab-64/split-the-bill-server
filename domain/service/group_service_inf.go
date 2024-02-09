package service

import (
	. "github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IGroupService interface {
	Create(groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error)

	Update(userID UUID, groupID UUID, group dto.GroupInput) (dto.GroupDetailedOutput, error)

	GetByID(id UUID) (dto.GroupDetailedOutput, error)

	GetAll(userID UUID, invitationID UUID) ([]dto.GroupDetailedOutput, error)

	AcceptGroupInvitation(invitationID UUID, userID UUID) error
}
