package storage_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

// UserStorage is a storage interface for users. Implementation must make sure that no two stored users can have the
// same ID or email.

type IUserStorage interface {
	// Delete deletes the user with the given ID from the storage, if it exists.
	Delete(id UUID) error

	// GetAll returns all users in the storage.
	GetAll() ([]UserModel, error)

	// GetByID returns the user with the given ID, or a NoSuchUserError if no such user exists.
	GetByID(id UUID) (UserModel, error)

	// GetByEmail returns the user with the given email, or a NoSuchUserError if no such user exists.
	GetByEmail(email string) (UserModel, error)

	// Create adds the given user to the storage and saves the password. If a user with the same ID or email
	// already exists, a UserAlreadyExistsError is returned.
	Create(user UserModel, passwordHash []byte) error

	// GetCredentials returns the password hash for the user with the given ID, or a NoCredentialsError, if no
	// credentials are stored for the user.
	GetCredentials(id UUID) ([]byte, error)
}
