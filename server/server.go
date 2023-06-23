package server

import "github.com/gofiber/fiber/v2"

// SetupRoutes creates webserver routes
//
// Parameters:
//
//	*fiber.App: The fiber server to be configured
func SetupRoutes(app *fiber.App) {
	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
