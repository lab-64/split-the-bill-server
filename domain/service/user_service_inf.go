package service

import (
	"github.com/google/uuid"
	"mime/multipart"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
)

type IUserService interface {
	// Delete deleted the user with the given id.
	// *Authorization required: requesterID == id
	Delete(requesterID uuid.UUID, id uuid.UUID) error

	GetAll() ([]dto.UserCoreOutput, error)

	GetByID(id uuid.UUID) (dto.UserCoreOutput, error)

	Login(userInput dto.UserInput) (dto.UserCoreOutput, model.AuthCookie, error)

	// Logout logs out the user with the given token.
	// *Authorization required: requesterID == cookie.UserID
	Logout(requesterID uuid.UUID, token uuid.UUID) error

	Create(user dto.UserInput) (dto.UserCoreOutput, error)

	// Update updates the user with the given id with the new user data.
	// *Authorization required: requesterID == id
	Update(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate, profileImg multipart.File) (dto.UserCoreOutput, error)
}
