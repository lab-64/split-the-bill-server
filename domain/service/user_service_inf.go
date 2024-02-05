package service

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	. "split-the-bill-server/presentation/dto"
)

type IUserService interface {
	Delete(id UUID) error

	GetAll() ([]UserDetailedOutputDTO, error)

	GetByID(id UUID) (UserDetailedOutputDTO, error)

	Login(credentials CredentialsInputDTO) (UserCoreOutputDTO, AuthCookieModel, error)

	Create(user UserInputDTO) (UserCoreOutputDTO, error)

	Update(requesterID UUID, id UUID, user UserUpdateDTO) (UserCoreOutputDTO, error)
}
