package storage

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type ICookieStorage interface {
	AddAuthenticationCookie(cookie types.AuthenticationCookie)

	GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie

	// GetCookieFromToken returns the authentication cookie for the given token, or a NoSuchCookieError if no such
	GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error)
}
