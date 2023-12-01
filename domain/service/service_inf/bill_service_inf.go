package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill BillInputDTO) (BillOutputDTO, error)

	GetByID(id UUID) (BillOutputDTO, error)

	AddItem(item ItemInputDTO) (ItemOutputDTO, error)
}
