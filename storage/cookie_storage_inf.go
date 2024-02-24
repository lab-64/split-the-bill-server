package storage

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type ICookieStorage interface {
	AddAuthenticationCookie(cookie model.AuthCookie) (model.AuthCookie, error)

	// Delete deletes the cookie with the given uuid. If an error occurs, a NoSuchCookieError is returned
	Delete(token UUID) error

	GetCookiesForUser(userID UUID) []model.AuthCookie

	// GetCookieFromToken returns the authentication cookie for the given token, or a NoSuchCookieError if no such
	GetCookieFromToken(token UUID) (model.AuthCookie, error)
}
