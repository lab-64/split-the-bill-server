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
	log.Println("running migrations")
	// migrate models to database
	err = db.AutoMigrate(&User{}, &AuthCookie{}, &Credentials{})
	if err != nil {
		return err
	}
	// set database
	d.db = db
	return nil
}

func (d *Database) AddUser(user types.User) error {
	item := MakeUser(user)
	// FIXME: This is a little bit of TOCTOU:
	// If a user with the same username is created after the check, we DO NOT RETURN AN ERROR.
	// We also do not overwrite the existing user.
	// Checking, if the username already exists is still better, as we will receive an error at least in most cased,
	// where there are no unlikely race condition.
	// This could be fixed, if there was some way to check, whether FirstOrCreate actually created a new user or not.
	_, err := d.GetUserByUsername(user.Username) // check if user already exists
	if err == nil {
		return storage.UserAlreadyExistsError
	}
	res := d.db.Where(User{Username: user.Username}).FirstOrCreate(&item) // write new user if not exists
	return res.Error
}

func (d *Database) DeleteUser(id uuid.UUID) error {
	tx := d.db.Delete(&User{}, "id = ?", id)
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

func (d *Database) AddAuthenticationCookie(cookie types.AuthenticationCookie) {
	authCookie := MakeAuthCooke(cookie)
	// store cookie
	d.db.Where(AuthCookie{UserID: authCookie.UserID}).Create(&authCookie)
}

func (d *Database) GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie {
	var cookies []AuthCookie
	// get all cookies for given user
	res := d.db.Where(AuthCookie{UserID: userID}).Find(&cookies)
	if res.Error != nil {
		return nil
	}
	return cookiesToAuthCookies(cookies)
}

// CookiesToAuthCookies converts a slice of AuthCookie to a slice of types.AuthenticationCookie
func cookiesToAuthCookies(cookies []AuthCookie) []types.AuthenticationCookie {
	s := make([]types.AuthenticationCookie, len(cookies))
	for i, cookie := range cookies {
		s[i] = cookie.ToAuthCookie()
	}
	return s
}

func (d *Database) RegisterUser(user types.User, passwordHash []byte) error {
	item := MakeUser(user)
	// store user
	res := d.db.Where(User{Username: item.Username}).FirstOrCreate(&item)
	log.Println(res)
	if res.Error != nil {
		return storage.UserAlreadyExistsError
	}
	// TODO: handle error case, user should not be created, if credentials cannot be stored
	// store credentials
	res = d.db.Where(Credentials{UserID: item.ID}).FirstOrCreate(&Credentials{UserID: item.ID, Hash: passwordHash})
	if res.Error != nil {
		// TODO: create suitable error msg
		return res.Error
	}
	return nil
}

func (d *Database) GetCredentials(id uuid.UUID) ([]byte, error) {
	var credentials Credentials
	// get credentials from given user
	res := d.db.Limit(1).First(&credentials, Credentials{UserID: id})
	if res.Error != nil {
		return nil, storage.NoCredentialsError
	}
	return credentials.Hash, nil
}
