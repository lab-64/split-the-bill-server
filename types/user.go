package types

import "github.com/google/uuid"

type User struct {
	ID                   uuid.UUID `json:"id,omitempty"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	ConfirmationPassword string    `json:"confirmationPassword"`
}

func NewUser(username string) User {
	return User{
		ID:       uuid.New(),
		Username: username,
	}
}

func (u User) Equals(other User) bool {
	return u.ID == other.ID && u.Username == other.Username
}
