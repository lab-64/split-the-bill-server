package model

import "github.com/google/uuid"

type ItemModel struct {
	ID    uuid.UUID
	Name  string
	Price float64
}

func CreateItemModel(name string, price float64) ItemModel {
	return ItemModel{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	}
}
