package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage/storage_inf"
)

var (
	MockUserDelete         func(id uuid.UUID) error
	MockUserGetByID        func(id uuid.UUID) (model.UserModel, error)
	MockUserGetAll         func() ([]model.UserModel, error)
	MockUserGetByEmail     func(email string) (model.UserModel, error)
	MockUserCreate         func(user model.UserModel, passwordHash []byte) error
	MockUserGetCredentials func(id uuid.UUID) ([]byte, error)
)

func NewUserStorageMock() storage_inf.IUserStorage {
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

func (u UserStorageMock) Create(userModel model.UserModel, passwordHash []byte) error {
	return MockUserCreate(userModel, passwordHash)
}

func (u UserStorageMock) GetCredentials(id uuid.UUID) ([]byte, error) {
	return MockUserGetCredentials(id)
}
