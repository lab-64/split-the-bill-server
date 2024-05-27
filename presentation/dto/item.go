package dto

import (
	"errors"
	"github.com/google/uuid"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ItemInput struct {
	ID           uuid.UUID           `json:"id"`
	Name         *string             `json:"name,omitempty"`
	Price        *float64            `json:"price,omitempty"`
	Contributors *Changes[uuid.UUID] `json:"contributorIDs,omitempty"`
}

type ItemOutput struct {
	ID           uuid.UUID        `json:"id"`
	Name         string           `json:"name"`
	Price        float64          `json:"price"`
	BillID       uuid.UUID        `json:"billId"`
	Contributors []UserCoreOutput `json:"contributors"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validators
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (i ItemInput) ValidateInputs() error {
	if *i.Name == "" {
		return ErrItemNameRequired
	}
	if *i.Price == 0 {
		return ErrItemPriceRequired
	}
	return nil
}

var ErrItemNameRequired = errors.New("name is required")
var ErrItemPriceRequired = errors.New("price is required")
