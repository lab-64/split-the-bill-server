package storage_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type ICookieStorage interface {
	AddAuthenticationCookie(cookie model.AuthenticationCookie)

	GetCookiesForUser(userID uuid.UUID) []model.AuthenticationCookie

	// GetCookieFromToken returns the authentication cookie for the given token, or a NoSuchCookieError if no such
	GetCookieFromToken(token uuid.UUID) (model.AuthenticationCookie, error)
}
