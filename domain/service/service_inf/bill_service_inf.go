package service_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill dto.BillInputDTO) (dto.BillOutputDTO, error)

	GetByID(id uuid.UUID) (dto.BillOutputDTO, error)
}
