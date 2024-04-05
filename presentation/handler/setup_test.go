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
	userService  service.IUserService
	groupService service.IGroupService
	userHandler  UserHandler
	groupHandler GroupHandler
	app          *fiber.App
)

func Authenticate(c *fiber.Ctx) error {
	c.Locals(middleware.UserKey, TestUser.ID)
	return c.Next()
}

func TestMain(m *testing.M) {
	// setup mocks
	userService = mocks.NewUserServiceMock()
	groupService = mocks.NewGroupServiceMock()
	// password validator
	passwordValidator, err := util.NewPasswordValidator()
	if err != nil {
		panic("Error while setting up the password validator: " + err.Error())
	}
	// setup handlers
	userHandler = *NewUserHandler(&userService, passwordValidator)
	groupHandler = *NewGroupHandler(&groupService)

	// setup fiber
	app = fiber.New()
	app.Get("/api/user/:id", userHandler.GetByID)
	app.Post("/api/user/register", userHandler.Register)
	app.Put("/api/user/:id", Authenticate, userHandler.Update)
	app.Post("/api/group", Authenticate, groupHandler.Create)

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}
