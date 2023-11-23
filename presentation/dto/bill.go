package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type BillInputDTO struct {
	Owner uuid.UUID `json:"owner"`
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Group uuid.UUID `json:"group"`
}

type BillOutputDTO struct {
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Items []ItemDTO `json:"items"`
}

func ToBillModel(b BillInputDTO) (BillModel, error) {
	return CreateBill(b.Owner, b.Name, b.Date), nil
}

func ToBillDTO(bill BillModel) BillOutputDTO {
	itemsDTO := make([]ItemDTO, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(item)
	}

	return BillOutputDTO{
		Name:  bill.Name,
		Date:  bill.Date,
		Items: itemsDTO,
	}
}
