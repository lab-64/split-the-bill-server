package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"split-the-bill-server/authentication"
	"split-the-bill-server/config"
	"split-the-bill-server/handler"
	"split-the-bill-server/router"
	"split-the-bill-server/storage/database"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// configure webserver
	app := fiber.New()
	storage, err := database.NewDatabase()
	err = storage.Connect()
	//storage := ephemeral.NewEphemeral()
	passwordValidator, err := authentication.NewPasswordValidator()
	if err != nil {
		log.Fatal(err)
	}
	h := handler.NewHandler(storage, storage, passwordValidator)
	router.SetupRoutes(app, h)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
