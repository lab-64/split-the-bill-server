package storage_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

// UserStorage is a storage interface for users. Implementation must make sure that no two stored users can have the
// same ID or username.

type IUserStorage interface {
	// Create adds the given user to the storage. If a user with the same ID or username already exists, a
	// UserAlreadyExistsError is returned.
	Create(model.User) error

	// Delete deletes the user with the given ID from the storage, if it exists.
	Delete(id uuid.UUID) error

	// GetAll returns all users in the storage.
	GetAll() ([]model.User, error)

	// GetByID returns the user with the given ID, or a NoSuchUserError if no such user exists.
	GetByID(id uuid.UUID) (model.User, error)

	// GetByUsername returns the user with the given username, or a NoSuchUserError if no such user exists.
	GetByUsername(username string) (model.User, error)

	// Register adds the given user to the storage and saves the password. If a user with the same ID or username
	// already exists, a UserAlreadyExistsError is returned.
	Register(user model.User, passwordHash []byte) error

	// GetCredentials returns the password hash for the user with the given ID, or a NoCredentialsError, if no
	// credentials are stored for the user.
	GetCredentials(id uuid.UUID) ([]byte, error)

	// AddGroupInvitation adds the given group invitation to the pending group invitations from user. If the user does not exist, a NoSuchUserError is returned.
	AddGroupInvitation(invitation model.GroupInvitation, userID uuid.UUID) error

	// HandleInvitation handles the given invitation for the given user. Invitations can be accepted or declined. If the user does not exist, a NoSuchUserError is returned.
	// If the invitation does not exist, a InvitationNotFoundError is returned.
	// TODO: rethink invitation type, maybe delete and only look for UUID in all invitations, or split into different handlers.
	HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error
}
