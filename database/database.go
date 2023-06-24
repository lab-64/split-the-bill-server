package database

import (
	"fmt"
	"log"
	"os"
	"split-the-bill-server/config"
	"split-the-bill-server/model"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database instance
type Database struct {
	Db *gorm.DB
}

var DB Database

// Connect Database
func Connect() {

	// convert port string to int
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	// insert postgresql configuration
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// check for connection failures
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	// successful connected
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{})

	// set database
	DB = Database{
		Db: db,
	}

}
