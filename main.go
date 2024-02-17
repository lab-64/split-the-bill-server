package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "split-the-bill-server/docs"
	"split-the-bill-server/domain/service/impl"
	"split-the-bill-server/domain/util"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
	"split-the-bill-server/presentation/router"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/db_storages"
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/ephemeral/eph_storages"
)

// @title		Split The Bill API
// @version	1.0
// @host		localhost:8080
// @BasePath	/
func main() {
	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// create storage directory for uploaded images
	if err = os.MkdirAll(router.StorePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// configure webserver
	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024, // 4MB
	})

	// setup storage
	userStorage, groupStorage, cookieStorage, billStorage := setupStorage()

	// services
	userService := impl.NewUserService(&userStorage, &cookieStorage)
	groupService := impl.NewGroupService(&groupStorage)
	billService := impl.NewBillService(&billStorage, &groupStorage)

	// password validator
	passwordValidator, err := util.NewPasswordValidator()
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := handler.NewUserHandler(&userService, passwordValidator)
	groupHandler := handler.NewGroupHandler(&groupService)
	billHandler := handler.NewBillHandler(&billService, &groupService)

	// setup logger
	app.Use(logger.New())

	// authenticator
	authenticator := middleware.NewAuthenticator(&cookieStorage)

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
		return db_storages.NewUserStorage(db), db_storages.NewGroupStorage(db), db_storages.NewCookieStorage(db), db_storages.NewBillStorage(db)
	case "ephemeral":
		db, err := ephemeral.NewEphemeral()
		if err != nil {
			log.Fatal(err)
		}
		// TODO: add invitation storage in ephemeral
		return eph_storages.NewUserStorage(db), eph_storages.NewGroupStorage(db), eph_storages.NewCookieStorage(db), eph_storages.NewBillStorage(db)
	default:
		log.Fatalf("Unsupported storage type: %s", storageType)
		return nil, nil, nil, nil
	}
}
