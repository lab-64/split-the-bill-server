package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
)

var (
	MockUserDelete  func(requesterID uuid.UUID, id uuid.UUID) error
	MockUserGetAll  func() ([]dto.UserCoreOutput, error)
	MockUserGetByID func(id uuid.UUID) (dto.UserCoreOutput, error)
	MockUserLogin   func(credentials dto.UserInput) (dto.UserCoreOutput, model.AuthCookie, error)
	MockUserCreate  func(user dto.UserInput) (dto.UserCoreOutput, error)
	MockUserUpdate  func(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate, file []byte) (dto.UserCoreOutput, error)
	MockUserLogout  func(requesterID uuid.UUID, token uuid.UUID) error
)

func NewUserServiceMock() service.IUserService {
	return &UserServiceMock{}
}

type UserServiceMock struct {
}

func (u UserServiceMock) Delete(requesterID uuid.UUID, id uuid.UUID) error {
	return MockUserDelete(requesterID, id)
}

func (u UserServiceMock) GetAll() ([]dto.UserCoreOutput, error) {
	return MockUserGetAll()
}

func (u UserServiceMock) GetByID(id uuid.UUID) (dto.UserCoreOutput, error) {
	return MockUserGetByID(id)
}

func (u UserServiceMock) Login(credentials dto.UserInput) (dto.UserCoreOutput, model.AuthCookie, error) {
	return MockUserLogin(credentials)
}

func (u UserServiceMock) Logout(requesterID uuid.UUID, token uuid.UUID) error {
	return MockUserLogout(requesterID, token)
}

func (u UserServiceMock) Create(user dto.UserInput) (dto.UserCoreOutput, error) {
	return MockUserCreate(user)
}

func (u UserServiceMock) Update(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate, file []byte) (dto.UserCoreOutput, error) {
	return MockUserUpdate(requesterID, id, user, file)
}
