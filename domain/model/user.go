package model

import (
	"github.com/google/uuid"
)

type UserModel struct {
	ID                      uuid.UUID
	Email                   string
	PendingGroupInvitations []GroupInvitationModel
	Groups                  []GroupModel
	Items                   []ItemModel
}

func CreateUserModel(email string) UserModel {
	return UserModel{
		ID:    uuid.New(),
		Email: email,
	}
}

func (u UserModel) Equals(other UserModel) bool {
	return u.ID == other.ID && u.Email == other.Email
}
