package router

import (
	"github.com/gofiber/fiber/v2"
	"split-the-bill-server/authentication"
	"split-the-bill-server/handler"
)

// SetupRoutes creates webserver routes and connect them to the related handlers.
func SetupRoutes(app *fiber.App, u handler.UserHandler, g handler.GroupHandler, b handler.BillHandler, a authentication.Authenticator) {

	// Define landing page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// grouping
	api := app.Group("/api")

	// user routes
	userRoute := api.Group("/user")

	// routes
	userRoute.Get("/", a.Authenticate, u.GetAll)
	userRoute.Get("/:id", a.Authenticate, u.GetByID)
	userRoute.Post("/", a.Authenticate, u.Create)
	userRoute.Get("/:username", a.Authenticate, u.GetByUsername)
	userRoute.Post("/register", u.Register)
	userRoute.Post("/login", u.Login)
	//userRoute.Put("/:id", u.UpdateUser)
	userRoute.Delete("/:id", a.Authenticate, u.Delete)
	userRoute.Post("/invitations", a.Authenticate, u.HandleInvitation)

	// bill routes
	billRoute := api.Group("/bill")
	// routes
	billRoute.Post("/", a.Authenticate, b.Create)
	billRoute.Get("/:id", a.Authenticate, b.GetByID)

	// group routes
	groupRoute := api.Group("/group")
	// routes
	groupRoute.Post("/", a.Authenticate, g.Create)
	groupRoute.Get("/:id", a.Authenticate, g.Get)
}
