package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	. "split-the-bill-server/storage/database/entity"
	"strconv"
)

type Database struct {
	Context *gorm.DB
}

func NewDatabase() (*Database, error) {
	d := Database{}
	err := d.Connect()
	return &d, err
}

func (d *Database) Connect() error {
	// convert port string to int
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Failed to parse port. \n", err)
	}

	// insert postgresql configuration
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)
	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// check for connection failures
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	// successful connected
	log.Printf("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// migrate models to database
	err = db.AutoMigrate(&User{}, &AuthCookie{}, &Credentials{}, &Group{}, &GroupInvitation{})
	if err != nil {
		return err
	}
	// set database
	d.Context = db
	return nil
}
