package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
)

type IUserService interface {
	// Delete deleted the user with the given id.
	// *Authorization required: requesterID == id
	Delete(requesterID uuid.UUID, id uuid.UUID) error

	GetAll() ([]dto.UserCoreOutput, error)

	GetByID(id uuid.UUID) (dto.UserCoreOutput, error)

	Login(credentials dto.CredentialsInput) (dto.UserCoreOutput, model.AuthCookie, error)

	Create(user dto.UserInput) (dto.UserCoreOutput, error)

	// Update updates the user with the given id with the new user data.
	// *Authorization required: requesterID == id
	Update(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate, profileImg []byte) (dto.UserCoreOutput, error)
}
