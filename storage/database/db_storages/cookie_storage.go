package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	. "split-the-bill-server/storage/storage_inf"
)

type CookieStorage struct {
	DB *gorm.DB
}

func NewCookieStorage(DB *Database) ICookieStorage {
	return &CookieStorage{DB: DB.Context}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie AuthCookieModel) {
	authCookie := ToAuthCookieEntity(cookie)
	// store cookie
	c.DB.Where(AuthCookie{UserID: authCookie.UserID}).Create(&authCookie)
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []AuthCookieModel {
	var cookies []AuthCookie
	// get all cookies for given user
	res := c.DB.Where(AuthCookie{UserID: userID}).Find(&cookies)
	if res.Error != nil {
		return nil
	}
	return cookiesToAuthCookies(cookies)
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (AuthCookieModel, error) {
	var cookie AuthCookie

	tx := c.DB.First(&cookie, token)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return AuthCookieModel{}, storage.NoSuchCookieError
		}
		return AuthCookieModel{}, tx.Error
	}

	return ToAuthCookieModel(&cookie), nil
}

// CookiesToAuthCookies converts a slice of AuthCookie to a slice of model.AuthCookieModel
func cookiesToAuthCookies(cookies []AuthCookie) []AuthCookieModel {
	s := make([]AuthCookieModel, len(cookies))
	for i, cookie := range cookies {
		s[i] = ToAuthCookieModel(&cookie)
	}
	return s
}
