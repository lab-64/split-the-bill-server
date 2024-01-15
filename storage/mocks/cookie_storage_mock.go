package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage/storage_inf"
)

var (
	MockCookieAddAuthenticationCookie func(cookie model.AuthCookieModel)
	MockCookieGetCookiesForUser       func(userID uuid.UUID) []model.AuthCookieModel
	MockCookieGetCookieFromToken      func(token uuid.UUID) (model.AuthCookieModel, error)
)

func NewCookieStorageMock() storage_inf.ICookieStorage {
	return &CookieStorageMock{}
}

type CookieStorageMock struct {
}

func (c CookieStorageMock) AddAuthenticationCookie(cookie model.AuthCookieModel) {
	MockCookieAddAuthenticationCookie(cookie)
}

func (c CookieStorageMock) GetCookiesForUser(userID uuid.UUID) []model.AuthCookieModel {
	return MockCookieGetCookiesForUser(userID)
}

func (c CookieStorageMock) GetCookieFromToken(token uuid.UUID) (model.AuthCookieModel, error) {
	return MockCookieGetCookieFromToken(token)
}
