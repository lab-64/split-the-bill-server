package wire

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

type Item struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributors"`
}

type Bill struct {
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Items *[]Item   `json:"items"`
	Group uuid.UUID `json:"group"`
}

// ToItem converts an Item to a types.Item. Returns an error if the conversion fails.
func (i Item) ToItem() (types.Item, error) {
	return types.CreateItem(i.Name, i.Price, i.Contributors), nil
}

// ToBill converts a Bill to a types.Bill. Returns an error if the conversion fails.
func (b Bill) ToBill(owner uuid.UUID) (types.Bill, error) {
	// convert each item
	var items []types.Item
	for _, item := range *b.Items {
		convertedItem, err := item.ToItem()
		if err != nil {
			return types.Bill{}, err
		}
		items = append(items, convertedItem)
	}
	return types.CreateBill(owner, b.Name, b.Date, items), nil
}
