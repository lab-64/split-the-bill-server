package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockCookieAddAuthenticationCookie func(cookie model.AuthCookie) (model.AuthCookie, error)
	MockCookieGetCookiesForUser       func(userID uuid.UUID) []model.AuthCookie
	MockCookieGetCookieFromToken      func(token uuid.UUID) (model.AuthCookie, error)
	MockCookieDelete                  func(token uuid.UUID) error
)

func NewCookieStorageMock() storage.ICookieStorage {
	return &CookieStorageMock{}
}

type CookieStorageMock struct {
}

func (c CookieStorageMock) AddAuthenticationCookie(cookie model.AuthCookie) (model.AuthCookie, error) {
	return MockCookieAddAuthenticationCookie(cookie)
}

func (c CookieStorageMock) GetCookiesForUser(userID uuid.UUID) []model.AuthCookie {
	return MockCookieGetCookiesForUser(userID)
}

func (c CookieStorageMock) GetCookieFromToken(token uuid.UUID) (model.AuthCookie, error) {
	return MockCookieGetCookieFromToken(token)
}

func (c CookieStorageMock) Delete(token uuid.UUID) error {
	return MockCookieDelete(token)
}
