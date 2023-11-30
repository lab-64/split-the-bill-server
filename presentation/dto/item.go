package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type ItemInputDTO struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	BillID       uuid.UUID   `json:"billId"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

type ItemContributorInputDTO struct {
	ItemID       uuid.UUID   `json:"itemId"`
	Contributors []uuid.UUID `json:"contributors"`
}

type ItemOutputDTO struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

// ToItemModel converts an ItemInputDTO to an ItemModel
func ToItemModel(i ItemInputDTO) ItemModel {
	return CreateItemModel(i.Name, i.Price, i.BillID, i.Contributors)
}

// ToItemDTO converts an ItemModel to an ItemOutputDTO
func ToItemDTO(item ItemModel) ItemOutputDTO {
	return ItemOutputDTO{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		Contributors: item.Contributors,
	}
}
