package types

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username"`
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
