package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func (d *Database) LoginUser(userCredentials types.AuthenticateCredentials) (types.AuthCookie, error) {
	// check if user exist with given credentials
	var user User
	// get user
	res := d.db.Limit(1).First(&user, User{Email: userCredentials.Email})
	if res.Error != nil {
		return types.AuthCookie{}, res.Error
	}
	// compare passwords
	err := ComparePasswords(user, userCredentials.Password)
	if err != nil {
		return types.AuthCookie{}, err
	}
	// check if cookie already exists
	cookie, err := d.GetCookieFromUser(user.ID)
	if err != nil {
		// no cookie -> create cookie
		cookie, err = d.CreateAuthCookie(user.ID)
		if err != nil {
			// TODO: Change error msg
			return types.AuthCookie{}, err
		}
	}

	return cookie, nil
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

func (d *Database) CreateAuthCookie(userId uuid.UUID) (types.AuthCookie, error) {
	var cookie = MakeAuthCookie(userId)
	// tries to create a new AuthCookie
	res := d.db.Where(AuthCookie{Base: Base{ID: cookie.ID}}).FirstOrCreate(&cookie)
	return cookie.ToCookie(), res.Error
}

func (d *Database) GetUserFromAuthCookie(cookieId uuid.UUID) (types.User, error) {
	// search cookie
	var cookie AuthCookie
	dbRes := d.db.Limit(1).First(&cookie, AuthCookie{Base: Base{ID: cookieId}})
	if dbRes.Error != nil {
		return types.User{}, dbRes.Error
	}
	// search user from cookie
	var user User
	dbRes = d.db.Limit(1).First(&user, User{Base: Base{ID: cookie.UserId}})
	if dbRes.Error != nil {
		return types.User{}, dbRes.Error
	}

	return user.ToUser(), nil
}

func (d *Database) GetCookieFromUser(userId uuid.UUID) (types.AuthCookie, error) {
	var cookie AuthCookie
	res := d.db.Limit(1).First(&cookie, AuthCookie{UserId: userId})
	if res.Error != nil {
		return types.AuthCookie{}, res.Error
	}
	return cookie.ToCookie(), nil
}
