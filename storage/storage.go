package storage

import (
	"errors"
	"split-the-bill-server/types"

	"github.com/google/uuid"
)

// TODO: Add generic storage tests

type Storage interface {
	// Connect connects to the storage and must be called exactly once before interacting with the storage.
	Connect() error
}

// UserStorage is a storage interface for users. Implementation must make sure that no two stored users can have the
// same ID or username.
type UserStorage interface {
	Storage
	// AddUser adds the given user to the storage. If a user with the same ID or username already exists, a
	// UserAlreadyExistsError is returned.
	AddUser(types.User) (types.User, error)
	// LoginUser creates and returns an authentication token if the given credentials are valid otherwise an error.
	LoginUser(types.AuthenticateCredentials) (types.AuthCookie, error)
	// DeleteUser deletes the user with the given ID from the storage, if it exists.
	DeleteUser(id uuid.UUID) error
	// GetAllUsers returns all users in the storage.
	GetAllUsers() ([]types.User, error)
	// GetUserByID returns the user with the given ID, or a NoSuchUserError if no such user exists.
	GetUserByID(id uuid.UUID) (types.User, error)
	// GetUserByUsername returns the user with the given username, or a NoSuchUserError if no such user exists.
	GetUserByUsername(username string) (types.User, error)
}

type CookieStorage interface {
	Storage
	// CreateAuthCookie creates an authentication cookie for the given user ID.
	CreateAuthCookie(userId uuid.UUID) (types.AuthCookie, error)
	// Get user from authentication cookie
	GetUserFromAuthCookie(cookieId uuid.UUID) (types.User, error)
	// GetCookieFromUser return a valid (non-expired) cookie from the given user if such an cookie exists. Otherwise it return an error.
	GetCookieFromUser(userId uuid.UUID) (types.AuthCookie, error)
}

type AuthenticatedUserStorage interface {
	UserStorage
	CookieStorage
}

var UserAlreadyExistsError = errors.New("user already exists")
var NoSuchUserError = errors.New("no such user")
