package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

type ItemDTO struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributors"`
}

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

// ToItem converts an ItemDTO to a types.Item. Returns an error if the conversion fails.
func (i ItemDTO) ToItem() (types.Item, error) {
	return types.CreateItem(i.Name, i.Price, i.Contributors), nil
}

// ToBill converts a BillInputDTO to a types.Bill. Returns an error if the conversion fails.
func (b BillInputDTO) ToBill(owner uuid.UUID) (types.Bill, error) {
	// convert each item
	var items []types.Item
	for _, item := range b.Items {
		convertedItem, err := item.ToItem()
		if err != nil {
			return types.Bill{}, err
		}
		items = append(items, convertedItem)
	}
	return types.CreateBill(owner, b.Name, b.Date, items), nil
}

func ToBillDTO(bill *types.Bill) BillOutputDTO {
	itemsDTO := make([]ItemDTO, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(&item)
	}

	return BillOutputDTO{
		Name:  bill.Name,
		Date:  bill.Date,
		Items: itemsDTO,
	}
}

func ToItemDTO(item *types.Item) ItemDTO {
	return ItemDTO{
		Name:         item.Name,
		Price:        item.Price,
		Contributors: item.Contributors,
	}
}
