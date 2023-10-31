package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	storage "split-the-bill-server/storage/storage_inf"
)

type CookieStorage struct {
	DB *gorm.DB
}

func NewCookieStorage(DB *database.Database) storage.ICookieStorage {
	return &CookieStorage{DB: DB.Context}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie model.AuthenticationCookie) {
	authCookie := MakeAuthCooke(cookie)
	// store cookie
	c.DB.Where(AuthCookie{UserID: authCookie.UserID}).Create(&authCookie)
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []model.AuthenticationCookie {
	var cookies []AuthCookie
	// get all cookies for given user
	res := c.DB.Where(AuthCookie{UserID: userID}).Find(&cookies)
	if res.Error != nil {
		return nil
	}
	return cookiesToAuthCookies(cookies)
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (model.AuthenticationCookie, error) {
	//TODO implement me
	panic("implement me")
}

// CookiesToAuthCookies converts a slice of AuthCookie to a slice of model.AuthenticationCookie
func cookiesToAuthCookies(cookies []AuthCookie) []model.AuthenticationCookie {
	s := make([]model.AuthenticationCookie, len(cookies))
	for i, cookie := range cookies {
		s[i] = cookie.ToAuthCookie()
	}
	return s
}
