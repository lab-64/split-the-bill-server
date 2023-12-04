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

type ItemOutputDTO struct {
	ID           uuid.UUID   `json:"id"`
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	BillID       uuid.UUID   `json:"billId"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

func ToItemModel(id uuid.UUID, i ItemInputDTO) ItemModel {
	return CreateItemModel(id, i.Name, i.Price, i.Contributors, i.BillID)
}

// ToItemDTO converts an ItemModel to an ItemOutputDTO
func ToItemDTO(item ItemModel) ItemOutputDTO {
	return ItemOutputDTO{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: item.Contributors,
	}
}
