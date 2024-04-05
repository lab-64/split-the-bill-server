package model

import (
	"github.com/google/uuid"
)

type GroupInvitation struct {
	ID    uuid.UUID
	Group Group
}
