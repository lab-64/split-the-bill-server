package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type ICookieService interface {
	AddAuthenticationCookie(cookie types.AuthenticationCookie)

	GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie

	GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error)
}
