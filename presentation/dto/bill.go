package dto

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type BillCreate struct {
	OwnerID uuid.UUID   `json:"ownerID"`
	Name    string      `json:"name"`
	Date    time.Time   `json:"date"`
	GroupID uuid.UUID   `json:"groupID"`
	Items   []ItemInput `json:"items"`
}

type BillUpdate struct {
	Name   string      `json:"name"`
	Date   time.Time   `json:"date"`
	Viewed bool        `json:"isViewed,omitempty"`
	Items  []ItemInput `json:"items"`
}

type BillDetailedOutput struct {
	ID      uuid.UUID             `json:"id"`
	Name    string                `json:"name"`
	Date    time.Time             `json:"date"`
	Items   []ItemOutput          `json:"items"`
	GroupID uuid.UUID             `json:"groupID"`
	Owner   UserCoreOutput        `json:"owner"`
	Balance map[uuid.UUID]float64 `json:"balance,omitempty"` // include balance only if balance is set
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validators
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b BillCreate) ValidateInputs() error {
	if b.OwnerID == uuid.Nil {
		return ErrBillOwnerIDRequired
	}
	if b.Name == "" {
		return ErrBillNameRequired
	}
	if b.Date.IsZero() {
		return ErrBillDateRequired
	}
	if b.GroupID == uuid.Nil {
		return ErrBillGroupIDRequired
	}
	return nil
}

var ErrBillOwnerIDRequired = errors.New("ownerID is required")
var ErrBillNameRequired = errors.New("name is required")
var ErrBillDateRequired = errors.New("date is required")
var ErrBillGroupIDRequired = errors.New("groupID is required")
