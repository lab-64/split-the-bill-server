package model

import "github.com/google/uuid"

type ItemModel struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	BillID       uuid.UUID
	Contributors []UserModel
}

func CreateItemModel(id uuid.UUID, name string, price float64, contributors []UserModel, billID uuid.UUID) ItemModel {
	if id == uuid.Nil {
		id = uuid.New()
	}

	return ItemModel{
		ID:           id,
		Name:         name,
		Price:        price,
		Contributors: contributors,
		BillID:       billID,
	}
}
