package model

import (
	"github.com/google/uuid"
	"time"
)

type AuthenticationCookie struct {
	UserID      uuid.UUID
	Token       uuid.UUID
	ValidBefore time.Time
}
