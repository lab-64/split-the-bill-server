package router

import (
	"github.com/gofiber/fiber/v2"
	"os"
	. "split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
)

// SetupRoutes creates webserver routes and connect them to the related handlers.
func SetupRoutes(app *fiber.App, u UserHandler, g GroupHandler, b BillHandler, a middleware.Authenticator) {

	// Serve static files (images, css, js, etc.)
	app.Static("/", "presentation/pages/public")

	// Define pages
	app.Get("/", func(c *fiber.Ctx) error { // landing page
		return c.SendFile("presentation/pages/homepage.html")
	})
	app.Get("/privacy-policy", func(c *fiber.Ctx) error { // privacy policy
		return c.SendFile("presentation/pages/privacy_policy.html")
	})
	app.Get("/delete-user", func(c *fiber.Ctx) error { // delete account page
		return c.SendFile("presentation/pages/delete_account.html")
	})
	app.Get("/impressum", func(c *fiber.Ctx) error { // impressum
		return c.SendFile("presentation/pages/impressum.html")
	})

	// Manage DeepLink
	app.Get("/.well-known/assetlinks.json", func(c *fiber.Ctx) error {
		// Read the assetlinks.json file
		fileContent, err := os.ReadFile(".well-known/assetlinks.json")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to read assetlinks.json file")
		}
		c.Set("Content-Type", "application/json")
		// Serve the JSON response
		return c.Send(fileContent)
	})

	// serve static images to authenticated users
	app.Use("/image/", a.Authenticate)
	app.Static("/image/", "./uploads/profileImages") // hide the real storage path

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
	userRoute.Post("/logout", a.Authenticate, u.Logout)

	// bill routes
	billRoute := api.Group("/bill")
	// routes
	billRoute.Post("/", a.Authenticate, b.Create)
	billRoute.Put("/:id/contribution", a.Authenticate, b.UpdateContribution)
	billRoute.Put("/:id", a.Authenticate, b.Update)
	billRoute.Get("/:id", a.Authenticate, b.GetByID)
	billRoute.Delete("/:id", a.Authenticate, b.Delete)
	billRoute.Get("/", a.Authenticate, b.GetAllByUser)

	// group routes
	groupRoute := api.Group("/group")
	// routes
	groupRoute.Post("/invitation/:id/accept", a.Authenticate, g.AcceptInvitation)
	groupRoute.Post("/:id/transaction/", a.Authenticate, g.CreateGroupTransaction)
	groupRoute.Get("/transaction/", a.Authenticate, g.GetAllGroupTransactions)
	groupRoute.Put("/:id", a.Authenticate, g.Update)
	groupRoute.Get("/:id", a.Authenticate, g.GetByID)
	groupRoute.Delete("/:id", a.Authenticate, g.Delete)
	groupRoute.Post("/", a.Authenticate, g.Create)
	groupRoute.Get("/", a.Authenticate, g.GetAll)
}
