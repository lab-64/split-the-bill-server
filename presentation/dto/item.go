package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type BaseItem struct {
	Name           string      `json:"name"`
	Price          float64     `json:"price"`
	BillID         uuid.UUID   `json:"billId"`
	ContributorIDs []uuid.UUID `json:"contributorIDs"`
}

type Item struct {
	ID uuid.UUID `json:"id"`
	BaseItem
}

func ToItemModel(id uuid.UUID, item BaseItem) ItemModel {
	return CreateItemModel(id, item.Name, item.Price, item.ContributorIDs, item.BillID)
}

func ToItemDTO(item ItemModel) Item {
	return Item{
		ID: item.ID,
		BaseItem: BaseItem{
			Name:           item.Name,
			Price:          item.Price,
			BillID:         item.BillID,
			ContributorIDs: item.Contributors,
		},
	}
}
