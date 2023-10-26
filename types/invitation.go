package types

import (
	"github.com/google/uuid"
	"time"
)

type GroupInvitation struct {
	ID   uuid.UUID
	Date time.Time
	For  Group
}

func CreateGroupInvitation(group Group) GroupInvitation {
	return GroupInvitation{
		ID:   uuid.New(),
		Date: time.Now(),
		For:  group,
	}
}
