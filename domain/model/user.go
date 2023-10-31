package model

import (
	"github.com/google/uuid"
)

type UserModel struct {
	ID                      uuid.UUID
	Username                string
	Email                   string
	PendingGroupInvitations []GroupInvitationModel
	Groups                  []*GroupModel
}

func NewUser(username string) UserModel {
	return UserModel{
		ID:       uuid.New(),
		Username: username,
	}
}

func CreateUserModel(username string, email string) UserModel {
	return UserModel{
		ID:                      uuid.New(),
		Username:                username,
		Email:                   email,
		PendingGroupInvitations: make([]GroupInvitationModel, 0),
		Groups:                  make([]*GroupModel, 0),
	}
}

func (u UserModel) Equals(other UserModel) bool {
	return u.ID == other.ID && u.Username == other.Username
}
