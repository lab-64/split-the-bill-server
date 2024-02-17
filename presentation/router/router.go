package router

import (
	"github.com/gofiber/fiber/v2"
	. "split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
)

var UploadPath = "/image/"
var StorePath = "./uploads/profileImages"

// SetupRoutes creates webserver routes and connect them to the related handlers.
func SetupRoutes(app *fiber.App, u UserHandler, g GroupHandler, b BillHandler, a middleware.Authenticator) {

	// Define landing page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// serve static files to authenticated user
	app.Use(UploadPath, a.Authenticate)
	app.Static(UploadPath, "./uploads/profileImages") // hide the real storage path

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
	itemRoute.Delete("/:id", a.Authenticate, b.DeleteItem)

	// group routes
	groupRoute := api.Group("/group")
	// routes
	groupRoute.Post("/", a.Authenticate, g.Create)
	groupRoute.Put("/:id", a.Authenticate, g.Update)
	groupRoute.Get("/:id", a.Authenticate, g.GetByID)
	groupRoute.Get("/", a.Authenticate, g.GetAll)
	groupRoute.Delete("/:id", a.Authenticate, g.Delete)
	groupRoute.Post("/invitation/:id/accept", a.Authenticate, g.AcceptInvitation)
}
