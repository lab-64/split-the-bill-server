package authentication

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

const SessionCookieValidityPeriod = time.Hour * 24 * 7

func GenerateSessionCookie(userID uuid.UUID) types.AuthenticationCookie {
	// TODO: Safely generate a session cookie.
	return types.AuthenticationCookie{
		UserID:      userID,
		Token:       uuid.New(),
		ValidBefore: time.Now().Add(SessionCookieValidityPeriod),
	}
}
