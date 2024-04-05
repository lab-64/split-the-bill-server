package entity

import (
	"github.com/google/uuid"
	"time"
)

// AuthCookie struct
type AuthCookie struct {
	Base
	User        User
	UserID      uuid.UUID `gorm:"type:uuid; not null"`
	ValidBefore time.Time
}
