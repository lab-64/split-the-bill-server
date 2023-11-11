package service_inf

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IUserService interface {
	Delete(id UUID) error

	GetAll() ([]UserOutputDTO, error)

	GetByID(id UUID) (UserOutputDTO, error)

	GetByUsername(username string) (UserOutputDTO, error)

	Login(credentials CredentialsInputDTO) (fiber.Cookie, error)

	Register(user UserInputDTO) (UserOutputDTO, error)

	HandleInvitation(invitation InvitationInputDTO) error
}
