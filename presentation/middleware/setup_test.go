package middleware

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"split-the-bill-server/presentation"
	storagemocks "split-the-bill-server/storage/mocks"
	"testing"
)

var (
	app           *fiber.App
	authenticator *Authenticator
)

func TestMain(m *testing.M) {
	// setup authentication
	cookieStorage := storagemocks.NewCookieStorageMock()
	authenticator = NewAuthenticator(&cookieStorage)

	// setup fiber
	app = fiber.New()
	// setup test route
	app.Get("/user", authenticator.Authenticate, func(c *fiber.Ctx) error { return presentation.Success(c, fiber.StatusOK, "Authentication accept", nil) })

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}
