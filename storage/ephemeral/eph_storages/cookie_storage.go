package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	eph "split-the-bill-server/storage/ephemeral"
)

type CookieStorage struct {
	e *eph.Ephemeral
}

func NewCookieStorage(ephemeral *eph.Ephemeral) storage.ICookieStorage {
	return &CookieStorage{e: ephemeral}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie model.AuthCookie) (model.AuthCookie, error) {
	r := c.e.Locker.Lock(eph.RCookies)
	defer c.e.Locker.Unlock(r)
	cookies, exists := c.e.Cookies[cookie.UserID]
	if !exists {
		cookies = make([]model.AuthCookie, 0)
	}
	cookies = append(cookies, cookie)
	c.e.Cookies[cookie.UserID] = cookies
	return cookie, nil
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []model.AuthCookie {
	r := c.e.Locker.Lock(eph.RCookies)
	defer c.e.Locker.Unlock(r)
	return c.e.Cookies[userID]
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (model.AuthCookie, error) {
	r := c.e.Locker.Lock(eph.RCookies)
	defer c.e.Locker.Unlock(r)
	for _, cookies := range c.e.Cookies {
		for _, cookie := range cookies {
			if cookie.Token == token {
				return cookie, nil
			}
		}
	}
	return model.AuthCookie{}, storage.NoSuchCookieError
}

func (c *CookieStorage) Delete(token uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
