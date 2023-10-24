package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

// AuthCookie struct
type AuthCookie struct {
	Base
	UserID      uuid.UUID `gorm:"type:uuid; column:user_foreign_key;not null"`
	ValidBefore time.Time
}

func MakeAuthCooke(authCookie types.AuthenticationCookie) AuthCookie {
	return AuthCookie{Base: Base{ID: authCookie.Token}, UserID: authCookie.UserID, ValidBefore: authCookie.ValidBefore}
}

func (authCookie *AuthCookie) ToAuthCookie() types.AuthenticationCookie {
	return types.AuthenticationCookie{UserID: authCookie.UserID, Token: authCookie.ID, ValidBefore: authCookie.ValidBefore}
}
