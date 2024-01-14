package storage

import (
	"errors"
)

// User Errors
var UserAlreadyExistsError = errors.New("user already exists")
var NoSuchUserError = errors.New("no such user")

// Credentials Errors
var NoCredentialsError = errors.New("no credentials for user")
var NoSuchCookieError = errors.New("no such cookie")

// Group Errors
var GroupAlreadyExistsError = errors.New("group already exists")
var NoSuchGroupError = errors.New("no such group")

// Group Invitation Errors
var GroupInvitationAlreadyExistsError = errors.New("group invitation already exists")
var NoSuchGroupInvitationError = errors.New("no such group invitation")

// Bill Errors
var BillAlreadyExistsError = errors.New("bill already exists")
var NoSuchBillError = errors.New("no such bill")

// Item Errors
var ItemAlreadyExistsError = errors.New("item already exists")
var NoSuchItemError = errors.New("no such item")
