package service_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill BillInputDTO) (BillOutputDTO, error)

	GetByID(id UUID) (BillOutputDTO, error)

	AddItemToBill(billID UUID, item ItemCreateDTO) (ItemOutputDTO, error)

	ChangeItem(item ItemEditDTO) (ItemOutputDTO, error)

	GetItemByID(id UUID) (ItemOutputDTO, error)
}
