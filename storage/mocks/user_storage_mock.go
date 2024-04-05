package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockUserDelete         func(id uuid.UUID) error
	MockUserGetByID        func(id uuid.UUID) (model.User, error)
	MockUserGetAll         func() ([]model.User, error)
	MockUserGetByEmail     func(email string) (model.User, error)
	MockUserCreate         func(user model.User, passwordHash []byte) (model.User, error)
	MockUserGetCredentials func(id uuid.UUID) ([]byte, error)
	MockUserUpdate         func(user model.User) (model.User, error)
)

func NewUserStorageMock() storage.IUserStorage {
	return &UserStorageMock{}
}

type UserStorageMock struct {
}

func (u UserStorageMock) Delete(id uuid.UUID) error {
	return MockUserDelete(id)
}

func (u UserStorageMock) GetAll() ([]model.User, error) {
	return MockUserGetAll()
}

func (u UserStorageMock) GetByID(id uuid.UUID) (model.User, error) {
	return MockUserGetByID(id)
}

func (u UserStorageMock) GetByEmail(email string) (model.User, error) {
	return MockUserGetByEmail(email)
}

func (u UserStorageMock) Create(user model.User, passwordHash []byte) (model.User, error) {
	return MockUserCreate(user, passwordHash)
}

func (u UserStorageMock) GetCredentials(id uuid.UUID) ([]byte, error) {
	return MockUserGetCredentials(id)
}

func (u UserStorageMock) Update(user model.User) (model.User, error) {
	return MockUserUpdate(user)
}
