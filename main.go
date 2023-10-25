package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"log"
	"split-the-bill-server/authentication"
	"split-the-bill-server/config"
	_ "split-the-bill-server/docs"
	"split-the-bill-server/handler"
	"split-the-bill-server/router"
	"split-the-bill-server/service/impl"
	"split-the-bill-server/storage/ephemeral"
)

// @title		Split The Bill API
// @version	1.0
// @host		localhost:8080
// @BasePath	/
func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// configure webserver
	app := fiber.New()

	// Select storage
	// start postgres
	/*storage, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}*/

	// start ephemeral
	e, _ := ephemeral.NewEphemeral()

	//storages
	userStorage := ephemeral.NewUserStorage(e)
	groupStorage := ephemeral.NewGroupStorage(e)
	cookieStorage := ephemeral.NewCookieStorage(e)
	billStorage := ephemeral.NewBillStorage(e)

	//services
	userService := impl.NewUserService(&userStorage, &cookieStorage)
	groupService := impl.NewGroupService(&groupStorage, &userStorage)
	cookieService := impl.NewCookieService(&cookieStorage)
	billService := impl.NewBillService(&billStorage, &groupStorage)

	//password validator
	passwordValidator, err := authentication.NewPasswordValidator()
	if err != nil {
		log.Fatal(err)
	}

	//handlers
	userHandler := handler.NewUserHandler(&userService, &cookieService, passwordValidator)
	groupHandler := handler.NewGroupHandler(&userService, &groupService)
	billHandler := handler.NewBillHandler(&billService, &groupService)

	// setup logger
	app.Use(logger.New())

	//routing
	router.SetupRoutes(app, *userHandler, *groupHandler, *billHandler)

	// setup swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// setup CORS
	app.Use(cors.New())

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
