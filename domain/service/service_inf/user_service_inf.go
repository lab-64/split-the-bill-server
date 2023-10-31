package service_inf

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	. "split-the-bill-server/presentation/dto"
)

type IUserService interface {
	Create(user UserInputDTO) (UserOutputDTO, error)

	Delete(id UUID) error

	GetAll() ([]UserOutputDTO, error)

	GetByID(id UUID) (UserOutputDTO, error)

	GetByUsername(username string) (UserOutputDTO, error)

	Login(credentials CredentialsInputDTO) (fiber.Cookie, error)

	Register(user UserInputDTO) (UserOutputDTO, error)

	// AddGroupInvitation adds the given group invitation to user's pending invitations.
	AddGroupInvitation(invitation GroupInvitationModel, userID UUID) error

	HandleInvitation(invitation InvitationInputDTO, userID UUID, invitationID UUID) error

	GetAuthenticatedUserID(tokenUUID UUID) (UUID, error)
}
