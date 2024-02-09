package service

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
)

type IUserService interface {
	Delete(id UUID) error

	GetAll() ([]dto.UserCoreOutput, error)

	GetByID(id UUID) (dto.UserCoreOutput, error)

	Login(credentials dto.CredentialsInput) (dto.UserCoreOutput, model.AuthCookie, error)

	Create(user dto.UserInput) (dto.UserCoreOutput, error)

	Update(requesterID UUID, id UUID, user dto.UserUpdate) (dto.UserCoreOutput, error)
}
