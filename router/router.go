package router

import (
	"split-the-bill-server/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes creates webserver routes
//
// Parameters:
//
//	app: The fiber server to be configured
func SetupRoutes(app *fiber.App, h handler.Handler) {

	// Define landing page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// grouping
	api := app.Group("/api")

	// user routes
	userRoute := api.Group("/user")
	// routes
	//userRoute.Get("/", handler.GetAllUsers)
	//userRoute.Get("/:id", handler.GetSingleUser)
	userRoute.Get("/:username", h.CreateUser)
	//userRoute.Put("/:id", handler.UpdateUser)
	//userRoute.Delete("/:id", handler.DeleteUserByID)
}
