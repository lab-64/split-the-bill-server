package storage

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type ICookieStorage interface {
	AddAuthenticationCookie(cookie AuthCookieModel)

	GetCookiesForUser(userID UUID) []AuthCookieModel

	// GetCookieFromToken returns the authentication cookie for the given token, or a NoSuchCookieError if no such
	GetCookieFromToken(token UUID) (AuthCookieModel, error)
}
