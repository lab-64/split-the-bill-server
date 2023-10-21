package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/dto"
)

type IGroupService interface {
	Create(groupDTO dto.GroupCreateDTO) (dto.GroupDTO, error)

	GetByID(id uuid.UUID) (dto.GroupDTO, error)
}
