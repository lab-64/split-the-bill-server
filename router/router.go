package router

import (
	"github.com/gofiber/fiber/v2"
	"split-the-bill-server/handler"
)

// SetupRoutes creates webserver routes and connect them to the related handlers.
func SetupRoutes(app *fiber.App, u handler.UserHandler, g handler.GroupHandler, b handler.BillHandler, i handler.InvitationHandler) {

	// Define landing page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// grouping
	api := app.Group("/api")

	// user routes
	userRoute := api.Group("/user")

	// routes
	userRoute.Get("/", u.GetAll)
	userRoute.Get("/:id", u.GetByID)
	userRoute.Post("/", u.Create)
	userRoute.Get("/:username", u.GetByUsername)
	userRoute.Post("/register", u.Register)
	userRoute.Post("/login", u.Login)
	//userRoute.Put("/:id", u.UpdateUser)
	userRoute.Delete("/:id", u.Delete)
	userRoute.Post("/invitations", u.HandleInvitation)

	// bill routes
	billRoute := api.Group("/bill")
	// routes
	billRoute.Post("/", b.Create)
	billRoute.Get("/:id", b.GetByID)

	// group routes
	groupRoute := api.Group("/group")
	// routes
	groupRoute.Post("/", g.Create)
	groupRoute.Get("/:id", g.Get)

	// invitation routes
	invitationRoute := api.Group("/invitation")
	// routes
	invitationRoute.Post("/", i.Create)
	invitationRoute.Get("/:id", i.GetByID)
}
