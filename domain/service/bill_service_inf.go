package service

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill BillInputDTO) (BillDetailedOutputDTO, error)

	Update(userID UUID, billID UUID, billDTO BillInputDTO) (BillDetailedOutputDTO, error)

	GetByID(id UUID) (BillDetailedOutputDTO, error)

	GetAllByUserID(userID UUID, requesterID UUID) ([]BillDetailedOutputDTO, error)

	AddItem(item ItemInputDTO) (ItemOutputDTO, error)

	ChangeItem(itemID UUID, item ItemInputDTO) (ItemOutputDTO, error)

	GetItemByID(id UUID) (ItemOutputDTO, error)
}
