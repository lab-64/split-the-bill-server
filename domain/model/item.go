package model

import "github.com/google/uuid"

type ItemModel struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	Contributors []uuid.UUID
}

func CreateItemModel(name string, price float64, contributors []uuid.UUID) ItemModel {
	return ItemModel{
		ID:           uuid.New(),
		Name:         name,
		Price:        price,
		Contributors: contributors,
	}
}