package service

import (
	. "github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill dto.BillInputDTO) (dto.BillDetailedOutputDTO, error)

	Update(userID UUID, billID UUID, billDTO dto.BillInputDTO) (dto.BillDetailedOutputDTO, error)

	GetByID(id UUID) (dto.BillDetailedOutputDTO, error)

	AddItem(item dto.Item) (dto.Item, error)

	ChangeItem(itemID UUID, item dto.Item) (dto.Item, error)

	GetItemByID(id UUID) (dto.Item, error)
}
