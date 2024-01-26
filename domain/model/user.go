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
