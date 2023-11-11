package core

import (
	"github.com/joho/godotenv"
)

// LoadConfig loads the .env file.
func LoadConfig() error {
	// load .env file
	return godotenv.Load(".env")
}
