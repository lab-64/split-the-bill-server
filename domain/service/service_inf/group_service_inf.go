package service_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IGroupService interface {
	Create(groupDTO dto.GroupInputDTO) (dto.GroupOutputDTO, error)

	GetByID(id uuid.UUID) (dto.GroupOutputDTO, error)
}
