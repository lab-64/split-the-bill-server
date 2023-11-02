package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

// AuthCookie struct
type AuthCookie struct {
	Base
	User        User
	UserID      uuid.UUID `gorm:"type:uuid; not null"`
	ValidBefore time.Time
}

func ToAuthCookieEntity(authCookie AuthCookieModel) AuthCookie {
	return AuthCookie{
		Base:        Base{ID: authCookie.Token},
		UserID:      authCookie.UserID,
		ValidBefore: authCookie.ValidBefore,
	}
}

func ToAuthCookieModel(authCookie *AuthCookie) AuthCookieModel {
	return AuthCookieModel{
		UserID:      authCookie.UserID,
		Token:       authCookie.ID,
		ValidBefore: authCookie.ValidBefore,
	}
}
