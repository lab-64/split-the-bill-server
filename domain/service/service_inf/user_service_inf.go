package service_inf

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/google/uuid"
	. "split-the-bill-server/presentation/dto"
)

type IUserService interface {
	Delete(id UUID) error

	GetAll() ([]UserOutputDTO, error)

	GetDetailedDataByID(id UUID) (UserOutputDTO, error)

	Login(credentials CredentialsInputDTO) (UserOutputDTO, fiber.Cookie, error)

	Register(user UserInputDTO) (UserOutputDTO, error)
}
