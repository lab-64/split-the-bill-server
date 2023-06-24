package main

import (
	"log"

	"split-the-bill-server/database"

	"split-the-bill-server/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// connect database
	database.Connect()
	// configure webserver
	app := fiber.New()
	router.SetupRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
