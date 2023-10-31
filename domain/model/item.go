package model

import "github.com/google/uuid"

type Item struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	Contributors []uuid.UUID
}

func CreateItem(name string, price float64, contributors []uuid.UUID) Item {
	return Item{
		ID:           uuid.New(),
		Name:         name,
		Price:        price,
		Contributors: contributors,
	}
}
