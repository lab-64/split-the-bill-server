package test_util

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

// Users
var User = UserModel{ID: uuid.New(), Email: "tester@mail.com"}
var UserWithEmptyEmail = UserModel{ID: uuid.New()}
var UserWithSameEmail = UserModel{ID: uuid.New(), Email: "tester@mail.com"}
var UserWithSameID = UserModel{ID: User.ID, Email: "new@mail.com"}

// Credentials
var PasswordHash = []byte("hashed_password")