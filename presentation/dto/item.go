package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type ItemDTO struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributors"`
}

func ToItemModel(i ItemDTO) ItemModel {
	return CreateItemModel(i.Name, i.Price, i.Contributors)
}

func ToItemDTO(item ItemModel) ItemDTO {
	return ItemDTO{
		Name:         item.Name,
		Price:        item.Price,
		Contributors: item.Contributors,
	}
}
