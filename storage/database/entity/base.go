package entity

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}
