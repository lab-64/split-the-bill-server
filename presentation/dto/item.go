package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type ItemCreateDTO struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

type ItemEditDTO struct {
	ID           uuid.UUID   `json:"id"`
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

func ToItemModelCreate(i ItemCreateDTO) ItemModel {
	return CreateItemModel(uuid.Nil, i.Name, i.Price, i.Contributors, uuid.Nil)
}

func ToItemModelEdit(i ItemEditDTO) ItemModel {
	return CreateItemModel(i.ID, i.Name, i.Price, i.Contributors, i.BillID)
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
