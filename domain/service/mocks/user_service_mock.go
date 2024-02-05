package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
)

var (
	MockUserDelete  func(id uuid.UUID) error
	MockUserGetAll  func() ([]dto.UserDetailedOutputDTO, error)
	MockUserGetByID func(id uuid.UUID) (dto.UserDetailedOutputDTO, error)
	MockUserLogin   func(credentials dto.CredentialsInputDTO) (dto.UserCoreOutputDTO, model.AuthCookieModel, error)
	MockUserCreate  func(user dto.UserInputDTO) (dto.UserCoreOutputDTO, error)
	MockUserUpdate  func(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdateDTO) (dto.UserCoreOutputDTO, error)
)

func NewUserServiceMock() service.IUserService {
	return &UserServiceMock{}
}

type UserServiceMock struct {
}

func (u UserServiceMock) Delete(id uuid.UUID) error {
	return MockUserDelete(id)
}

func (u UserServiceMock) GetAll() ([]dto.UserDetailedOutputDTO, error) {
	return MockUserGetAll()
}

func (u UserServiceMock) GetByID(id uuid.UUID) (dto.UserDetailedOutputDTO, error) {
	return MockUserGetByID(id)
}

func (u UserServiceMock) Login(credentials dto.CredentialsInputDTO) (dto.UserCoreOutputDTO, model.AuthCookieModel, error) {
	return MockUserLogin(credentials)
}

func (u UserServiceMock) Create(user dto.UserInputDTO) (dto.UserCoreOutputDTO, error) {
	return MockUserCreate(user)
}

func (u UserServiceMock) Update(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdateDTO) (dto.UserCoreOutputDTO, error) {
	return MockUserUpdate(requesterID, id, user)
}
