package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
	"split-the-bill-server/types"
)

type IUserService interface {
	Create(user dto.UserInputDTO) (dto.UserOutputDTO, error)

	Delete(id uuid.UUID) error

	GetAll() ([]dto.UserOutputDTO, error)

	GetByID(id uuid.UUID) (dto.UserOutputDTO, error)

	GetByUsername(username string) (dto.UserOutputDTO, error)

	Login(credentials dto.CredentialsInputDTO) (fiber.Cookie, error)

	Register(user dto.UserInputDTO) (dto.UserOutputDTO, error)

	// AddGroupInvitation adds the given group invitation to user's pending invitations.
	AddGroupInvitation(invitation types.GroupInvitation, userID uuid.UUID) error

	HandleInvitation(invitation dto.InvitationInputDTO, userID uuid.UUID, invitationID uuid.UUID) error

	GetAuthenticatedUserID(tokenUUID uuid.UUID) (uuid.UUID, error)
}
