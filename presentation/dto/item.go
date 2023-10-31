package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type ItemDTO struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributors"`
}

func (i ItemDTO) ToItem() (model.Item, error) {
	return model.CreateItem(i.Name, i.Price, i.Contributors), nil
}

func ToItemDTO(item *model.Item) ItemDTO {
	return ItemDTO{
		Name:         item.Name,
		Price:        item.Price,
		Contributors: item.Contributors,
	}
}
