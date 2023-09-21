package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"split-the-bill-server/storage"
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
	log.Printf("running migrations")
	err = db.AutoMigrate(&User{}, &AuthCookie{})
	if err != nil {
		return err
	}
	// set database
	d.db = db
	return nil
}

func (d *Database) AddUser(user types.User) (types.User, error) {
	item, err := MakeUser(user)
	if err != nil {
		return types.User{}, err
	}
	// FIXME: This is a little bit of TOCTOU:
	// If a user with the same username is created after the check, we DO NOT RETURN AN ERROR.
	// We also do not overwrite the existing user.
	// Checking, if the username already exists is still better, as we will receive an error at least in most cased,
	// where there are no unlikely race condition.
	// This could be fixed, if there was some way to check, whether FirstOrCreate actually created a new user or not.
	res := d.db.Where(User{Email: user.Email}).FirstOrCreate(&item) // tries to create new user
	return item.ToUser(), res.Error
}

}

func (d *Database) DeleteUser(id uuid.UUID) error {
	tx := d.db.Unscoped().Delete(&User{}, "id = ?", id)
	return tx.Error
}

func (d *Database) GetAllUsers() ([]types.User, error) {
	var users []User
	// find all users in the database
	// TODO: GetAllUsers should not return an error, if no users are found
	tx := d.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// return users
	return ToUserSlice(users), nil
}

func (d *Database) GetUserByID(uuid uuid.UUID) (types.User, error) {
	var user User
	tx := d.db.Limit(1).Find(&user, "id = ?", uuid)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return types.User{}, storage.NoSuchUserError
		}
		return types.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return types.User{}, storage.NoSuchUserError
	}
	return user.ToUser(), nil
}

func (d *Database) GetUserByUsername(username string) (types.User, error) {
	var user User
	// TODO: Verify that this is in fact not injectable
	tx := d.db.Take(&user, "username = ?", username)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return types.User{}, storage.NoSuchUserError
		}
		return types.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return types.User{}, storage.NoSuchUserError
	}
	return user.ToUser(), nil
}
