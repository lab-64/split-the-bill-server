package storage_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

// UserStorage is a storage interface for users. Implementation must make sure that no two stored users can have the
// same ID or username.

type IUserStorage interface {
	// Delete deletes the user with the given ID from the storage, if it exists.
	Delete(id UUID) error

	// GetAll returns all users in the storage.
	GetAll() ([]UserModel, error)

	// GetByID returns the user with the given ID, or a NoSuchUserError if no such user exists.
	GetByID(id UUID) (UserModel, error)

	// GetByUsername returns the user with the given username, or a NoSuchUserError if no such user exists.
	GetByUsername(username string) (UserModel, error)

	// Create adds the given user to the storage and saves the password. If a user with the same ID or username
	// already exists, a UserAlreadyExistsError is returned.
	Create(user UserModel, passwordHash []byte) error

	// GetCredentials returns the password hash for the user with the given ID, or a NoCredentialsError, if no
	// credentials are stored for the user.
	GetCredentials(id UUID) ([]byte, error)

	// AddGroupInvitation adds the given group invitation to the pending group invitations from user. If the user does not exist, a NoSuchUserError is returned.
	AddGroupInvitation(invitation GroupInvitationModel, userID UUID) error

	// HandleInvitation handles the given invitation for the given user. Invitations can be accepted or declined. If the user does not exist, a NoSuchUserError is returned.
	// If the invitation does not exist, a InvitationNotFoundError is returned.
	// TODO: rethink invitation type, maybe delete and only look for UUID in all invitations, or split into different handlers.
	HandleInvitation(invitationType string, userID UUID, invitationID UUID, accept bool) error
}
