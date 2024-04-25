package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
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
)

// @title		Split The Bill API
// @version	1.0
// @host		stb-server3-jteqgnayba-ew.a.run.app
// @BasePath	/
func main() {

	// create storage directory for uploaded images
	if err := os.MkdirAll("./uploads/profileImages", os.ModePerm); err != nil {
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

	// load port, port is provided by Google Run
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

// setupStorage initializes and configures the storage components for a Google SQL database.
func setupStorage() (storage.IUserStorage, storage.IGroupStorage, storage.ICookieStorage, storage.IBillStorage) {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	return db_storages.NewUserStorage(db), db_storages.NewGroupStorage(db), db_storages.NewCookieStorage(db), db_storages.NewBillStorage(db)

}
