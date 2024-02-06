package model

import (
	"github.com/google/uuid"
)

type GroupInvitationModel struct {
	ID    uuid.UUID
	Group GroupModel
}
