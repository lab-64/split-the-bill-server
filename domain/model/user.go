package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID                      uuid.UUID
	Username                string
	Email                   string
	PendingGroupInvitations []GroupInvitation
	Groups                  []*Group
}

func NewUser(username string) User {
	return User{
		ID:       uuid.New(),
		Username: username,
	}
}

func CreateUser(username string, email string) User {
	return User{
		ID:                      uuid.New(),
		Username:                username,
		Email:                   email,
		PendingGroupInvitations: make([]GroupInvitation, 0),
		Groups:                  make([]*Group, 0),
	}
}

func (u User) Equals(other User) bool {
	return u.ID == other.ID && u.Username == other.Username
}
