package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill BillInputDTO) (BillDetailedOutputDTO, error)

	GetByID(id UUID) (BillDetailedOutputDTO, error)

	AddItem(item ItemInputDTO) (ItemOutputDTO, error)

	ChangeItem(itemID UUID, item ItemInputDTO) (ItemOutputDTO, error)

	GetItemByID(id UUID) (ItemOutputDTO, error)
}
