package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/service"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type CookieService struct {
	storage.ICookieStorage
}

func NewCookieService(cookieStorage *storage.ICookieStorage) service.ICookieService {
	return &CookieService{ICookieStorage: *cookieStorage}
}

func (c *CookieService) AddAuthenticationCookie(cookie types.AuthenticationCookie) {
	//TODO implement me
	panic("implement me")
}

func (c *CookieService) GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie {
	//TODO implement me
	panic("implement me")
}

func (c *CookieService) GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error) {
	//TODO implement me
	panic("implement me")
}
