package database

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/types"
)

type CookieStorage struct {
	DB *gorm.DB
}

func NewCookieStorage(DB *Database) storage.ICookieStorage {
	return &CookieStorage{DB: DB.context}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie types.AuthenticationCookie) {
	authCookie := MakeAuthCooke(cookie)
	// store cookie
	c.DB.Where(AuthCookie{UserID: authCookie.UserID}).Create(&authCookie)
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie {
	var cookies []AuthCookie
	// get all cookies for given user
	res := c.DB.Where(AuthCookie{UserID: userID}).Find(&cookies)
	if res.Error != nil {
		return nil
	}
	return cookiesToAuthCookies(cookies)
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error) {
	var cookie AuthCookie

	tx := c.DB.First(&cookie, token)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return types.AuthenticationCookie{}, gorm.ErrRecordNotFound
		}
		return types.AuthenticationCookie{}, tx.Error
	}

	return cookie.ToAuthCookie(), nil
}

// CookiesToAuthCookies converts a slice of AuthCookie to a slice of types.AuthenticationCookie
func cookiesToAuthCookies(cookies []AuthCookie) []types.AuthenticationCookie {
	s := make([]types.AuthenticationCookie, len(cookies))
	for i, cookie := range cookies {
		s[i] = cookie.ToAuthCookie()
	}
	return s
}
