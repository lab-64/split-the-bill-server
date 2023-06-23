package main

import (
	"log"
	"os"
	"split-the-bill-server/server"

	"github.com/gofiber/fiber/v2"
)

func SetLogFile(path string) {
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
}

func main() {
	app := fiber.New()

	// Set up routes
	server.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
