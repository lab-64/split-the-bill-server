package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type CookieStorage struct {
	e *Ephemeral
}

func NewCookieStorage(ephemeral *Ephemeral) storage.ICookieStorage {
	return &CookieStorage{e: ephemeral}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie types.AuthenticationCookie) {
	c.e.lock.Lock()
	defer c.e.lock.Unlock()
	cookies, exists := c.e.cookies[cookie.UserID]
	if !exists {
		cookies = make([]types.AuthenticationCookie, 0)
	}
	cookies = append(cookies, cookie)
	c.e.cookies[cookie.UserID] = cookies
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie {
	c.e.lock.Lock()
	defer c.e.lock.Unlock()
	return c.e.cookies[userID]
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error) {
	c.e.lock.Lock()
	defer c.e.lock.Unlock()
	for _, cookies := range c.e.cookies {
		for _, cookie := range cookies {
			if cookie.Token == token {
				return cookie, nil
			}
		}
	}
	return types.AuthenticationCookie{}, storage.NoSuchCookieError
}
