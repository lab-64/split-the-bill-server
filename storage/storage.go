package storage

import "split-the-bill-server/types"

type Storage interface {
	Connect() error
}

type UserStorage interface {
	Storage
	AddUser(types.User) error
	DeleteUser(types.User) error
	GetAllUsers() ([]types.User, error)
}
