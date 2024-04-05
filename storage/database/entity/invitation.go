package entity

import (
	"github.com/google/uuid"
)

type GroupInvitation struct {
	Base
	GroupID uuid.UUID `gorm:"type:uuid"`
}
