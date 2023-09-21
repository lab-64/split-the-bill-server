package types

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID                   uuid.UUID `json:"id,omitempty"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	ConfirmationPassword string    `json:"confirmationPassword"`
}

type AuthenticateCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(username string) User {
	return User{
		ID:       uuid.New(),
		Username: username,
	}
}

// ValidateInputs checks if input fields are left out. Returns an error if fields are not filled out, otherwise nil
func (c AuthenticateCredentials) ValidateInputs() error {
	err := errors.New("")
	// Check input fields an concate error
	if c.Password == "" {
		passwdErr := errors.New("Password is not filled out!")
		err = errors.Join(passwdErr)
	}
	if c.Email == "" {
		emailErr := errors.New("Email address is not filled out!")
		err = errors.Join(emailErr, err)
	}
	// Test if error is overwritten
	if err.Error() == "" {
		return nil
	}
	return err
}

func (u User) Equals(other User) bool {
	return u.ID == other.ID && u.Username == other.Username
}
