package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill BillInputDTO) (BillOutputDTO, error)

	GetByID(id UUID) (BillOutputDTO, error)

	AddItem(billID UUID, item ItemInputDTO) (ItemOutputDTO, error)

	ChangeItem(itemID UUID, billID UUID, item ItemInputDTO) (ItemOutputDTO, error)

	GetItemByID(id UUID) (ItemOutputDTO, error)
}
