package dto

import (
	"github.com/google/uuid"
)

type ItemInput struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	BillID       uuid.UUID   `json:"billId"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

type ItemOutput struct {
	ID           uuid.UUID        `json:"id"`
	Name         string           `json:"name"`
	Price        float64          `json:"price"`
	BillID       uuid.UUID        `json:"billId"`
	Contributors []UserCoreOutput `json:"contributors"`
}
