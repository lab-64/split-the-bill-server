package router

import (
	"github.com/gofiber/fiber/v2"
	. "split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
)

// SetupRoutes creates webserver routes and connect them to the related handlers.
func SetupRoutes(app *fiber.App, u UserHandler, g GroupHandler, b BillHandler, i InvitationHandler, a middleware.Authenticator) {

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
	userRoute.Post("/", u.Register)
	userRoute.Post("/login", u.Login)
	userRoute.Put("/:id", a.Authenticate, u.Update)
	userRoute.Delete("/:id", a.Authenticate, u.Delete)

	// bill routes
	billRoute := api.Group("/bill")
	// routes
	billRoute.Post("/", a.Authenticate, b.Create)
	billRoute.Put("/:id", a.Authenticate, b.Update)
	billRoute.Get("/:id", a.Authenticate, b.GetByID)
	// item routes
	itemRoute := billRoute.Group("/item")
	// routes
	itemRoute.Get("/:id", a.Authenticate, b.GetItemByID)
	itemRoute.Post("/", a.Authenticate, b.AddItem)
	itemRoute.Put("/:id", a.Authenticate, b.ChangeItem)

	// group routes
	groupRoute := api.Group("/group")
	// routes
	groupRoute.Post("/", a.Authenticate, g.Create)
	groupRoute.Put("/:id", a.Authenticate, g.Update)
	groupRoute.Get("/:id", a.Authenticate, g.GetByID)
	groupRoute.Get("/", a.Authenticate, g.GetAllByUser)

	// invitation routes
	invitationRoute := api.Group("/invitation")
	// routes
	// TODO: delete route
	invitationRoute.Post("/", i.Create)
	invitationRoute.Get("/:id", i.GetByID)
	invitationRoute.Post("/:id/accept", a.Authenticate, i.AcceptInvitation)

}
