package dto

import (
	"errors"
	"github.com/google/uuid"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type GroupInput struct {
	OwnerID uuid.UUID `json:"ownerID"`
	Name    string    `json:"name"`
}

type GroupDetailedOutput struct {
	Owner        UserCoreOutput        `json:"owner"`
	ID           uuid.UUID             `json:"id"`
	Name         string                `json:"name"`
	Members      []UserCoreOutput      `json:"members"`
	Bills        []BillDetailedOutput  `json:"bills"`
	Balance      map[uuid.UUID]float64 `json:"balance,omitempty"`      // include balance only if balance is set
	InvitationID uuid.UUID             `json:"invitationID,omitempty"` // include invitationID only if invitationID is set
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validators
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (g GroupInput) ValidateInput() error {
	if g.OwnerID == uuid.Nil {
		return ErrOwnerIDRequired

	}
	if g.Name == "" {
		return ErrGroupNameRequired
	}
	return nil
}

var ErrGroupNameRequired = errors.New("name is required")
var ErrOwnerIDRequired = errors.New("ownerID is required")
