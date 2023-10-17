package storage

import (
	"errors"
	"github.com/google/uuid"
	"split-the-bill-server/types"
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
	AddUser(types.User) error
	// DeleteUser deletes the user with the given ID from the storage, if it exists.
	DeleteUser(id uuid.UUID) error
	// GetAllUsers returns all users in the storage.
	GetAllUsers() ([]types.User, error)
	// GetUserByID returns the user with the given ID, or a NoSuchUserError if no such user exists.
	GetUserByID(id uuid.UUID) (types.User, error)
	// GetUserByUsername returns the user with the given username, or a NoSuchUserError if no such user exists.
	GetUserByUsername(username string) (types.User, error)
	// RegisterUser adds the given user to the storage and saves the password. If a user with the same ID or username
	// already exists, a UserAlreadyExistsError is returned.
	RegisterUser(user types.User, passwordHash []byte) error
	// GetCredentials returns the password hash for the user with the given ID, or a NoCredentialsError, if no
	// credentials are stored for the user.
	GetCredentials(id uuid.UUID) ([]byte, error)
	// AddGroupInvitationToUser adds the given group invitation to the pending group invitations from user. If the user does not exist, a NoSuchUserError is returned.
	AddGroupInvitationToUser(invitation types.GroupInvitation, user uuid.UUID) error
	// HandleInvitation handles the given invitation for the given user. Invitations can be accepted or declined. If the user does not exist, a NoSuchUserError is returned.
	// If the invitation does not exist, a InvitationNotFoundError is returned.
	// TODO: rethink invitation type, maybe delete and only look for UUID in all invitations, or split into different handlers.
	HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error
}

type CookieStorage interface {
	Storage
	AddAuthenticationCookie(cookie types.AuthenticationCookie)
	GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie
	// GetCookieFromToken returns the authentication cookie for the given token, or a NoSuchCookieError if no such
	GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error)
}

type GroupStorage interface {
	Storage
	// AddGroup adds the given group to the storage. If a group with the same ID or name already exists, a GroupAlreadyExistsError is returned.
	AddGroup(group types.Group) error
	// GetGroupByID returns the group with the given ID, or a NoSuchGroupError if no such group exists.
	GetGroupByID(id uuid.UUID) (types.Group, error)
	// AddMemberToGroup adds the given member to the group with the given ID. If the group does not exist, a NoSuchGroupError is returned.
	AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error
}

var InvitationNotFoundError = errors.New("invitation not found")
var UserAlreadyExistsError = errors.New("user already exists")
var NoSuchUserError = errors.New("no such user")
var NoCredentialsError = errors.New("no credentials for user")
var NoSuchCookieError = errors.New("no such cookie")
var GroupAlreadyExistsError = errors.New("group already exists")
var NoSuchGroupError = errors.New("no such group")
