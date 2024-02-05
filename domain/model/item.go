package model

import "github.com/google/uuid"

type ItemModel struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	BillID       uuid.UUID
	Contributors []UserModel
}
