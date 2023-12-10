package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"split-the-bill-server/storage/database/seed"
	"strconv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Failed to parse port. \n", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Running seeds")
	for _, s := range seed.All() {
		if err := s.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", s.Name, err)
		}
	}
	log.Println("Seeds run successfully")
}
