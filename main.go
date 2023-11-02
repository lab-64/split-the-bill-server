package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"log"
	"os"
	"split-the-bill-server/authentication"
	"split-the-bill-server/config"
	_ "split-the-bill-server/docs"
	"split-the-bill-server/handler"
	"split-the-bill-server/router"
	"split-the-bill-server/service/impl"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
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

	// setup storage
	userStorage, groupStorage, cookieStorage, billStorage := setupStorage()

	// services
	userService := impl.NewUserService(&userStorage, &cookieStorage)
	groupService := impl.NewGroupService(&groupStorage, &userStorage)
	billService := impl.NewBillService(&billStorage, &groupStorage)

	// password validator
	passwordValidator, err := authentication.NewPasswordValidator()
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := handler.NewUserHandler(&userService, passwordValidator)
	groupHandler := handler.NewGroupHandler(&userService, &groupService)
	billHandler := handler.NewBillHandler(&billService, &groupService)

	// setup logger
	app.Use(logger.New())

	// authenticator
	authenticator := authentication.NewAuthenticator(&cookieStorage)

	// routing
	router.SetupRoutes(app, *userHandler, *groupHandler, *billHandler, *authenticator)

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

// setupStorage initializes and configures the storage components based on the STORAGE_TYPE environment variable.
func setupStorage() (storage.IUserStorage, storage.IGroupStorage, storage.ICookieStorage, storage.IBillStorage) {

	storageType := os.Getenv("STORAGE_TYPE")

	switch storageType {
	case "postgres":
		db, err := database.NewDatabase()
		if err != nil {
			log.Fatal(err)
		}
		return database.NewUserStorage(db), database.NewGroupStorage(db), database.NewCookieStorage(db), database.NewBillStorage(db)
	case "ephemeral":
		db, err := ephemeral.NewEphemeral()
		if err != nil {
			log.Fatal(err)
		}
		return ephemeral.NewUserStorage(db), ephemeral.NewGroupStorage(db), ephemeral.NewCookieStorage(db), ephemeral.NewBillStorage(db)
	default:
		log.Fatalf("Unsupported storage type: %s", storageType)
		return nil, nil, nil, nil
	}
}
