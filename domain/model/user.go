package model

import (
	"github.com/google/uuid"
)

type UserModel struct {
	ID                      uuid.UUID
	Email                   string
	Username                string
	PendingGroupInvitations []GroupInvitationModel
	Groups                  []GroupModel
	Items                   []ItemModel
}

func CreateUser(id uuid.UUID, email string, username string) UserModel {
	return UserModel{
		ID:       id,
		Email:    email,
		Username: username,
	}
}
