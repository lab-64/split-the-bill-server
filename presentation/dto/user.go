package dto

import (
	"errors"
	"github.com/google/uuid"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Username string `json:"username"`
}

type UserCoreOutput struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validators
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (u UserInput) ValidateInputs() error {
	if u.Email == "" {
		return ErrEmailRequired
	}
	if u.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

var ErrEmailRequired = errors.New("email is required")
var ErrPasswordRequired = errors.New("password is required")
