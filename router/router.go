package router

import (
	"split-the-bill-server/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes creates webserver routes and connect them to the related handlers.
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
	userRoute.Get("/", h.GetAllUsers)
	userRoute.Get("/:id", h.GetUserByID)
	userRoute.Post("/", h.CreateUser)
	userRoute.Get("/:username", h.GetUserByUsername)
	//userRoute.Put("/:id", handler.UpdateUser)
	userRoute.Delete("/:id", h.DeleteUserByID)

	userRoute.Post("/register", h.RegisterUser)
}
