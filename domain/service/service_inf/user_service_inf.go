package service_inf

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IUserService interface {
	Delete(id UUID) error

	GetAll() ([]UserDetailedOutputDTO, error)

	GetByID(id UUID) (UserDetailedOutputDTO, error)

	Login(credentials CredentialsInputDTO) (UserCoreOutputDTO, fiber.Cookie, error)

	Register(user UserInputDTO) (UserCoreOutputDTO, error)
}
