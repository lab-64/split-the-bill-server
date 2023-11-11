package model

import (
	"github.com/google/uuid"
	"time"
)

type AuthCookieModel struct {
	UserID      uuid.UUID
	Token       uuid.UUID
	ValidBefore time.Time
}
