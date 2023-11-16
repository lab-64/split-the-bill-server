package storage

import (
	"errors"
)

// TODO: Add generic storage tests

type Connection interface {
	// Connect connects to the storage and must be called exactly once before interacting with the storage.
	Connect() error
}

var InvitationNotFoundError = errors.New("invitation not found")
var UserAlreadyExistsError = errors.New("user already exists")
var NoSuchUserError = errors.New("no such user")
var NoCredentialsError = errors.New("no credentials for user")
var NoSuchCookieError = errors.New("no such cookie")
var GroupAlreadyExistsError = errors.New("group already exists")
var NoSuchGroupError = errors.New("no such group")
var BillAlreadyExistsError = errors.New("bill already exists")
var NoSuchBillError = errors.New("no such bill")
var GroupInvitationAlreadyExistsError = errors.New("group invitation already exists")
var NoSuchGroupInvitationError = errors.New("no such group invitation")
