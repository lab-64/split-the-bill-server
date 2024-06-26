package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/converter"
	"split-the-bill-server/storage/database/entity"
)

type CookieStorage struct {
	DB *gorm.DB
}

func NewCookieStorage(DB *Database) ICookieStorage {
	return &CookieStorage{DB: DB.Context}
}

func (c *CookieStorage) AddAuthenticationCookie(cookie model.AuthCookie) (model.AuthCookie, error) {
	authCookie := converter.ToAuthCookieEntity(cookie)
	var storedCookie entity.AuthCookie
	// store cookie
	res := c.DB.Where(entity.AuthCookie{UserID: authCookie.UserID}).Assign(authCookie).FirstOrCreate(&storedCookie)
	if res.Error != nil {
		return model.AuthCookie{}, storage.UnexpectedError
	}
	return converter.ToAuthCookieModel(&storedCookie), nil
}

func (c *CookieStorage) Delete(token uuid.UUID) error {
	// delete cookie
	res := c.DB.Delete(&entity.AuthCookie{}, token)
	if res.Error != nil {
		return storage.NoSuchCookieError
	}
	return nil
}

func (c *CookieStorage) GetCookiesForUser(userID uuid.UUID) []model.AuthCookie {
	var cookies []entity.AuthCookie
	// get all cookies for given user
	res := c.DB.Where(entity.AuthCookie{UserID: userID}).Find(&cookies)
	if res.Error != nil {
		return nil
	}
	return converter.ToAuthCookieModels(cookies)
}

func (c *CookieStorage) GetCookieFromToken(token uuid.UUID) (model.AuthCookie, error) {
	var cookie entity.AuthCookie

	tx := c.DB.First(&cookie, token)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.AuthCookie{}, storage.NoSuchCookieError
		}
		return model.AuthCookie{}, tx.Error
	}

	return converter.ToAuthCookieModel(&cookie), nil
}
