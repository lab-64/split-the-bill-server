package test_util

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

// Users
var User = UserModel{ID: uuid.New(), Email: "tester@mail.com"}
var User2 = UserModel{ID: uuid.New(), Email: "tester2@mail.com"}
var UserWithEmptyEmail = UserModel{ID: uuid.New()}
var UserWithSameEmail = UserModel{ID: uuid.New(), Email: "tester@mail.com"}
var UserWithSameID = UserModel{ID: User.ID, Email: "new@mail.com"}

// Credentials
var Password = "test1337"
var PasswordHash = []byte("hashed_password")

// Cookies
var UserCookie = AuthCookieModel{
	UserID:      User.ID,
	Token:       uuid.New(),
	ValidBefore: time.Now().Add(time.Hour * 24 * 7),
}
