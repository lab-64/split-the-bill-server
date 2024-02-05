package handler

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/domain/service/mocks"
	"split-the-bill-server/domain/util"
	"split-the-bill-server/presentation/middleware"
	"testing"
)

var (
	userService service.IUserService
	userHandler UserHandler
	app         *fiber.App
)

func Authenticate(c *fiber.Ctx) error {
	c.Locals(middleware.UserKey, TestUser.ID)
	return c.Next()
}

func TestMain(m *testing.M) {
	// setup user handler
	userService = mocks.NewUserServiceMock()
	// password validator
	passwordValidator, err := util.NewPasswordValidator()
	if err != nil {
		panic("Error while setting up the password validator: " + err.Error())
	}
	userHandler = *NewUserHandler(&userService, passwordValidator)

	// setup fiber
	app = fiber.New()
	app.Get("/api/user/:id", userHandler.GetByID)
	app.Post("/api/user/register", userHandler.Register)
	app.Put("/api/user/:id", Authenticate, userHandler.Update)

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}
