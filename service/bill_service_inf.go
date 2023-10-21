package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/dto"
)

type IBillService interface {
	Create(bill dto.BillCreateDTO) (dto.BillDTO, error)

	GetByID(id uuid.UUID) (dto.BillDTO, error)
}
