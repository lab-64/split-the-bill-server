package types

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
}

func NewUser(username string) User {
	return User{
		ID:       uuid.New(),
		Username: username,
	}
}
