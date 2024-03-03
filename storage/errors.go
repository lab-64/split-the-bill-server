package storage

import (
	"errors"
)

// General Errors
var UnexpectedError = errors.New("an unexpected error occurred while interacting with the storage system. Please try again later or contact support if the issue persists")

// User Errors
var UserAlreadyExistsError = errors.New("user already exists")
var NoSuchUserError = errors.New("no such user")
var InvalidUserInputError = errors.New("invalid user input")

// Credentials Errors
var NoCredentialsError = errors.New("no credentials for user")
var NoSuchCookieError = errors.New("no such cookie")

// Group Errors
var GroupAlreadyExistsError = errors.New("group already exists")
var NoSuchGroupError = errors.New("no such group")

// Bill Errors
var BillAlreadyExistsError = errors.New("bill already exists")
var NoSuchBillError = errors.New("no such bill")

// Item Errors
var ItemAlreadyExistsError = errors.New("item already exists")
var NoSuchItemError = errors.New("no such item")
