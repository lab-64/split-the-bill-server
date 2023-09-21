package types

import (
	"time"

	"github.com/google/uuid"
)

type AuthCookie struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserID    uuid.UUID `json:"username"`
	ExpiredAt time.Time `json:"expirationTime"`
}
