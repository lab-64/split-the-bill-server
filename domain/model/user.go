package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Username string
}

func CreateUser(id uuid.UUID, email string, username string) User {
	return User{
		ID:       id,
		Email:    email,
		Username: username,
	}
}
