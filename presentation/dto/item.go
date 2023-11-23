package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type ItemInputDTO struct {
	Name   string    `json:"name"`
	Price  float64   `json:"price"`
	BillID uuid.UUID `json:"billId"`
}

type ItemOutputDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ToItemModel converts an ItemInputDTO to an ItemModel
func ToItemModel(i ItemInputDTO) ItemModel {
	return CreateItemModel(i.Name, i.Price, i.BillID)
}

// ToItemDTO converts an ItemModel to an ItemOutputDTO
func ToItemDTO(item ItemModel) ItemOutputDTO {
	return ItemOutputDTO{
		Name:  item.Name,
		Price: item.Price,
	}
}
