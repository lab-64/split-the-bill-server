package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockUserDelete         func(id uuid.UUID) error
	MockUserGetByID        func(id uuid.UUID) (model.UserModel, error)
	MockUserGetAll         func() ([]model.UserModel, error)
	MockUserGetByEmail     func(email string) (model.UserModel, error)
	MockUserCreate         func(user model.UserModel, passwordHash []byte) (model.UserModel, error)
	MockUserGetCredentials func(id uuid.UUID) ([]byte, error)
	MockUserUpdate         func(user model.UserModel) (model.UserModel, error)
)

func NewUserStorageMock() storage.IUserStorage {
	return &UserStorageMock{}
}

type UserStorageMock struct {
}

func (u UserStorageMock) Delete(id uuid.UUID) error {
	return MockUserDelete(id)
}

func (u UserStorageMock) GetAll() ([]model.UserModel, error) {
	return MockUserGetAll()
}

func (u UserStorageMock) GetByID(id uuid.UUID) (model.UserModel, error) {
	return MockUserGetByID(id)
}

func (u UserStorageMock) GetByEmail(email string) (model.UserModel, error) {
	return MockUserGetByEmail(email)
}

func (u UserStorageMock) Create(user model.UserModel, passwordHash []byte) (model.UserModel, error) {
	return MockUserCreate(user, passwordHash)
}

func (u UserStorageMock) GetCredentials(id uuid.UUID) ([]byte, error) {
	return MockUserGetCredentials(id)
}

func (u UserStorageMock) Update(user model.UserModel) (model.UserModel, error) {
	return MockUserUpdate(user)
}
