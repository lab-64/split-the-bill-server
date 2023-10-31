package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"time"
)

// AuthCookie struct
type AuthCookie struct {
	Base
	UserID      uuid.UUID `gorm:"type:uuid; column:user_foreign_key;not null"`
	ValidBefore time.Time
}

func MakeAuthCooke(authCookie model.AuthenticationCookie) AuthCookie {
	return AuthCookie{Base: Base{ID: authCookie.Token}, UserID: authCookie.UserID, ValidBefore: authCookie.ValidBefore}
}

func (authCookie *AuthCookie) ToAuthCookie() model.AuthenticationCookie {
	return model.AuthenticationCookie{UserID: authCookie.UserID, Token: authCookie.ID, ValidBefore: authCookie.ValidBefore}
}
