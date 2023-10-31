package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"time"
)

type BillOutputDTO struct {
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Items []ItemDTO `json:"items"`
}

type BillInputDTO struct {
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Items []ItemDTO `json:"items"`
	Group uuid.UUID `json:"group"`
}

func (b BillInputDTO) ToBill(owner uuid.UUID) (model.Bill, error) {
	// convert each item
	var items []*model.Item
	for _, item := range b.Items {
		convertedItem, err := item.ToItem()
		if err != nil {
			return model.Bill{}, err
		}
		items = append(items, &convertedItem)
	}
	return model.CreateBill(owner, b.Name, b.Date, items), nil
}

func ToBillDTO(bill *model.Bill) BillOutputDTO {
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
