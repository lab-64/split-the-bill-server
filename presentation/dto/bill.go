package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type BillInputDTO struct {
	Owner uuid.UUID       `json:"ownerID"`
	Name  string          `json:"name"`
	Date  time.Time       `json:"date"`
	Group uuid.UUID       `json:"groupID"`
	Items []ItemCreateDTO `json:"items"`
}

type BillOutputDTO struct {
	ID      uuid.UUID       `json:"id"`
	Name    string          `json:"name"`
	Date    time.Time       `json:"date"`
	Items   []ItemOutputDTO `json:"items"`
	OwnerID uuid.UUID       `json:"ownerID"`
}

// ToBillModel converts a BillInputDTO to a BillModel
func ToBillModel(b BillInputDTO) BillModel {
	// convert each item
	var items []ItemModel
	for _, item := range b.Items {
		items = append(items, ToItemModelCreate(item))
	}
	return CreateBillModel(b.Owner, b.Name, b.Date, b.Group, items)
}

// ToBillDTO converts a BillModel to a BillOutputDTO
func ToBillDTO(bill BillModel) BillOutputDTO {
	itemsDTO := make([]ItemOutputDTO, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(item)
	}

	return BillOutputDTO{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Items:   itemsDTO,
		OwnerID: bill.OwnerID,
	}
}
