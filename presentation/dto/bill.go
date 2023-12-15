package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type BillInputDTO struct {
	Owner uuid.UUID      `json:"ownerID"`
	Name  string         `json:"name"`
	Date  time.Time      `json:"date"`
	Group uuid.UUID      `json:"groupID"`
	Items []ItemInputDTO `json:"items"`
}

type BillCoreOutputDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type BillDetailedOutputDTO struct {
	ID      uuid.UUID       `json:"id"`
	Name    string          `json:"name"`
	Date    time.Time       `json:"date"`
	Items   []ItemOutputDTO `json:"items"`
	OwnerID uuid.UUID       `json:"ownerID"`
}

func ToBillModel(b BillInputDTO) BillModel {
	// convert each item
	var items []ItemModel
	for _, item := range b.Items {
		items = append(items, ToItemModel(uuid.Nil, item))
	}
	return CreateBillModel(b.Owner, b.Name, b.Date, b.Group, items)
}

func ToBillDetailedDTOs(bills []BillModel) []BillDetailedOutputDTO {
	billsDTO := make([]BillDetailedOutputDTO, len(bills))

	for i, bill := range bills {
		billsDTO[i] = ToBillDetailedDTO(bill)
	}
	return billsDTO
}

func ToBillDetailedDTO(bill BillModel) BillDetailedOutputDTO {
	itemsDTO := make([]ItemOutputDTO, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(item)
	}

	return BillDetailedOutputDTO{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Items:   itemsDTO,
		OwnerID: bill.OwnerID,
	}
}
