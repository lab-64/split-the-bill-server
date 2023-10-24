package authentication

import (
	"errors"
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

const SessionCookieValidityPeriod = time.Hour * 24 * 7
const SessionCookieName = "session_cookie"

func GenerateSessionCookie(userID uuid.UUID) types.AuthenticationCookie {
	// TODO: Safely generate a session cookie.
	return types.AuthenticationCookie{
		UserID:      userID,
		Token:       uuid.New(),
		ValidBefore: time.Now().Add(SessionCookieValidityPeriod),
	}
}

// IsSessionCookieValid validates the given session cookie by checking if the ValidBefore time is in the future. Returns SessionExpiredError if the cookie is not valid anymore.
func IsSessionCookieValid(cookie types.AuthenticationCookie) error {
	if cookie.ValidBefore.After(time.Now()) {
		return nil
	} else {
		return SessionExpiredError
	}
}

var SessionExpiredError = errors.New("session expired")
