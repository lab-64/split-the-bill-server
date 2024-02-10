package service

import (
	. "github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IBillService interface {
	Create(bill dto.BillInput) (dto.BillDetailedOutput, error)

	Update(userID UUID, billID UUID, billDTO dto.BillInput) (dto.BillDetailedOutput, error)

	GetByID(id UUID) (dto.BillDetailedOutput, error)

	AddItem(item dto.ItemInput) (dto.ItemOutput, error)

	ChangeItem(itemID UUID, item dto.ItemInput) (dto.ItemOutput, error)

	GetItemByID(id UUID) (dto.ItemOutput, error)
}
