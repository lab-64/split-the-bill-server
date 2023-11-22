package dto

import (
	. "split-the-bill-server/domain/model"
)

type ItemDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ToItemModel(i ItemDTO) ItemModel {
	return CreateItemModel(i.Name, i.Price)
}

func ToItemDTO(item ItemModel) ItemDTO {
	return ItemDTO{
		Name:  item.Name,
		Price: item.Price,
	}
}
