package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
)

type CookieStorage struct {
	e *ephemeral.Ephemeral
}

func NewCookieStorage(ephemeral *ephemeral.Ephemeral) storage.ICookieStorage {
	return &CookieStorage{e: ephemeral}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie model.AuthCookieModel) {
	c.e.Lock.Lock()
	defer c.e.Lock.Unlock()
	cookies, exists := c.e.Cookies[cookie.UserID]
	if !exists {
		cookies = make([]model.AuthCookieModel, 0)
	}
	cookies = append(cookies, cookie)
	c.e.Cookies[cookie.UserID] = cookies
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []model.AuthCookieModel {
	c.e.Lock.Lock()
	defer c.e.Lock.Unlock()
	return c.e.Cookies[userID]
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (model.AuthCookieModel, error) {
	c.e.Lock.Lock()
	defer c.e.Lock.Unlock()
	for _, cookies := range c.e.Cookies {
		for _, cookie := range cookies {
			if cookie.Token == token {
				return cookie, nil
			}
		}
	}
	return model.AuthCookieModel{}, storage.NoSuchCookieError
}
