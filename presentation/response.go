package presentation

import (
	"github.com/gofiber/fiber/v2"
	"split-the-bill-server/presentation/dto"
)

func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(dto.GeneralResponse{
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(dto.GeneralResponse{
		Message: message,
		Data:    nil,
	})
}
