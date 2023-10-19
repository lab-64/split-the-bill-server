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
	userRoute.Post("/register", h.RegisterUser)
	userRoute.Post("/login", h.Login)
	//userRoute.Put("/:id", handler.UpdateUser)
	userRoute.Delete("/:id", h.DeleteUserByID)
	userRoute.Post("/invitations", h.HandleInvitation)

	// bill routes
	billRoute := api.Group("/bill")
	// routes
	billRoute.Post("/create", h.CreateBill)
	billRoute.Get("/:id", h.GetBillByID)

	// group routes
	groupRoute := api.Group("/group")
	// routes
	groupRoute.Post("/create", h.CreateGroup)
	groupRoute.Get("/:id", h.GetGroup)
}
