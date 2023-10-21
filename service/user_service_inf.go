package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/dto"
)

type IUserService interface {
	Create(user dto.UserCreateDTO) (dto.UserDTO, error)

	Delete(id uuid.UUID) error

	GetAll() ([]dto.UserDTO, error)

	GetByID(id uuid.UUID) (dto.UserDTO, error)

	GetByUsername(username string) (dto.UserDTO, error)

	Login(c *fiber.Ctx, credentials dto.CredentialsDTO) error

	Register(user dto.UserCreateDTO) (dto.UserDTO, error)

	HandleInvitation(invitation dto.InvitationReplyDTO, userID uuid.UUID, invitationID uuid.UUID) error

	GetAuthenticatedUserID(tokenUUID uuid.UUID) (uuid.UUID, error)
}
