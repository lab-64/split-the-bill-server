package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type BillInputDTO struct {
	OwnerID uuid.UUID      `json:"ownerID"`
	Name    string         `json:"name"`
	Date    time.Time      `json:"date"`
	GroupID uuid.UUID      `json:"groupID"`
	Items   []ItemInputDTO `json:"items"`
}

type BillDetailedOutputDTO struct {
	ID      uuid.UUID         `json:"id"`
	Name    string            `json:"name"`
	Date    time.Time         `json:"date"`
	Items   []ItemOutputDTO   `json:"items"`
	GroupID uuid.UUID         `json:"groupID"`
	Owner   UserCoreOutputDTO `json:"owner"`
}

func CreateBillModel(id uuid.UUID, b BillInputDTO) BillModel {
	// convert each item
	var items []ItemModel
	for _, item := range b.Items {
		items = append(items, CreateItemModel(uuid.New(), item))
	}
	return BillModel{
		ID:      id,
		Owner:   UserModel{ID: b.OwnerID},
		Name:    b.Name,
		Date:    b.Date,
		GroupID: b.GroupID,
		Items:   items,
	}
}

func ConvertToBillDetailedDTOs(bills []BillModel) []BillDetailedOutputDTO {
	billsDTO := make([]BillDetailedOutputDTO, len(bills))

	for i, bill := range bills {
		billsDTO[i] = ConvertToBillDetailedDTO(bill)
	}
	return billsDTO
}

func ConvertToBillDetailedDTO(bill BillModel) BillDetailedOutputDTO {
	itemsDTO := make([]ItemOutputDTO, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ConvertToItemDTO(item)
	}

	return BillDetailedOutputDTO{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Items:   itemsDTO,
		Owner:   ConvertToUserCoreDTO(&bill.Owner),
		GroupID: bill.GroupID,
	}
}
