package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IGroupService interface {
	Create(groupDTO GroupInputDTO) (GroupOutputDTO, error)

	GetByID(id UUID) (GroupOutputDTO, error)
}
