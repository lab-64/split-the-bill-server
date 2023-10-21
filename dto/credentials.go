package dto

import "errors"

type CredentialsDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ValidateInputs checks if input fields are left out. Returns an error if fields are not filled out, otherwise nil
func (c CredentialsDTO) ValidateInputs() error {
	// TODO: Make sure that username may not be empty on register
	if c.Password == "" || c.Username == "" {
		return errors.New("username or password not filled out")
	}
	return nil
}
