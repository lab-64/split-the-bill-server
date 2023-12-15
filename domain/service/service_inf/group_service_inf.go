package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IGroupService interface {
	Create(groupDTO GroupInputDTO) (GroupDetailedOutputDTO, error)

	GetByID(id UUID) (GroupDetailedOutputDTO, error)

	GetAllByUser(userID UUID) ([]GroupDetailedOutputDTO, error)
}
