package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"split-the-bill-server/config"
	"split-the-bill-server/types"
	"strconv"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	d := Database{}
	err := d.Connect()
	return &d, err
}

func (d Database) Connect() error {

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
	}
	// successful connected
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	err = db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	// set database
	d.db = db
	return nil
}

func (d Database) AddUser(user types.User) error {
	item := MakeUser(user)
	err := d.db.Create(&item).Error
	if err != nil {
		return err
	}
	return nil
}

func (d Database) DeleteUser(user types.User) error {
	//TODO implement me
	panic("implement me")
}

func (d Database) GetAllUsers() ([]types.User, error) {
	//TODO implement me
	panic("implement me")
}
