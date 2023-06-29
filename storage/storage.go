package storage

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

// TODO: Add generic storage tests

type Storage interface {
	Connect() error
}

type UserStorage interface {
	Storage
	AddUser(types.User) error
	DeleteUser(id uuid.UUID) error
	GetAllUsers() ([]types.User, error)
	GetUserByID(id uuid.UUID) (types.User, error)
	GetUserByUsername(username string) (types.User, error)
}
