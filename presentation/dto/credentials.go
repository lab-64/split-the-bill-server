package dto

import "errors"

type CredentialsInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ValidateInputs checks if input fields are left out. Returns an error if fields are not filled out, otherwise nil
func (c CredentialsInput) ValidateInputs() error {
	if c.Password == "" || c.Email == "" {
		return errors.New("email or password not filled out")
	}
	return nil
}
