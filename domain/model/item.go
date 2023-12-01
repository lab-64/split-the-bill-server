package model

import "github.com/google/uuid"

type ItemModel struct {
	ID     uuid.UUID
	Name   string
	Price  float64
	BillID uuid.UUID
}

func CreateItemModel(name string, price float64, billID uuid.UUID) ItemModel {
	return ItemModel{
		ID:     uuid.New(),
		Name:   name,
		Price:  price,
		BillID: billID,
	}
}
