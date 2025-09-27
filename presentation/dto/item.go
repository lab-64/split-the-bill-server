package dto

import (
	"errors"
	"github.com/google/uuid"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ItemInput struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

type ContributionInput struct {
	Contribution []HasContributed `json:"contribution"`
}

type HasContributed struct {
	ItemID      uuid.UUID `json:"itemID"`
	Contributed bool      `json:"contributed"`
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
	if i.Name == "" {
		return ErrItemNameRequired
	}
	if i.Price == 0 {
		return ErrItemPriceRequired
	}
	return nil
}

var ErrItemNameRequired = errors.New("name is required")
var ErrItemPriceRequired = errors.New("price is required")
