package types

import (
	"github.com/google/uuid"
)

// TODO: Change PendingGroupInvitations to a pointer
type User struct {
	ID                      uuid.UUID         `json:"id,omitempty"`
	Username                string            `json:"username"`
	Email                   string            `json:"email"`
	PendingGroupInvitations []GroupInvitation `json:"pending-group-invitations"`
	Groups                  []*Group          `json:"groups"`
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
