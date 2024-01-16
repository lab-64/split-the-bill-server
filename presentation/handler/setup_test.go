package handler

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/domain/service/mocks"
	"split-the-bill-server/domain/util"
	"testing"
)

var (
	userService service.IUserService
	userHandler UserHandler
	app         *fiber.App
)

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
	app.Get("/user/:id", userHandler.GetByID)
	app.Post("/api/user/register", userHandler.Register)

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}
