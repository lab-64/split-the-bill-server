package model

import (
	"github.com/google/uuid"
	"time"
)

const SessionCookieValidityPeriod = time.Hour * 24 * 7

type AuthCookie struct {
	UserID      uuid.UUID
	Token       uuid.UUID
	ValidBefore time.Time
}

func GenerateSessionCookie(userID uuid.UUID) AuthCookie {
	// TODO: Safely generate a session cookie.
	return AuthCookie{
		UserID:      userID,
		Token:       uuid.New(),
		ValidBefore: time.Now().Add(SessionCookieValidityPeriod),
	}
}
