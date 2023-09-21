package database

import (
	"split-the-bill-server/types"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// User struct
type User struct {
	Base
	Email    string `gorm:"unique;not_null"`
	Password []byte `gorm:"not_null"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// Authentication Cookie struct
type AuthCookie struct {
	Base
	UserId    uuid.UUID `gorm:"type:uuid;column:user_foreign_key;not null;"`
	ExpiredAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

// Create a database user. Password gets hashed.
func MakeUser(user types.User) (User, error) {

	// Hash Password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return User{}, err
	}
	return User{Base: Base{ID: user.ID}, Email: user.Email, Password: password}, err
}

}

func (user *User) ToUser() types.User {
	return types.User{ID: user.ID, Email: user.Email, Password: string(user.Password), ConfirmationPassword: string(user.Password)}
}

func ToUserSlice(users []User) []types.User {
	s := make([]types.User, len(users))
	for i, user := range users {
		s[i] = user.ToUser()
	}
	return s
}

// ToCookie converts a model.AuthCookie object to a types.AuthCookie object.
func (authCookie *AuthCookie) ToCookie() types.AuthCookie {
	return types.AuthCookie{ID: authCookie.ID, UserID: authCookie.UserId, ExpiredAt: authCookie.ExpiredAt}
}

// MakeAuthCookie creates an authentication cookie.
func MakeAuthCookie(userId uuid.UUID) AuthCookie {
	// Cookie available 14 Days
	var expirationTime = time.Now().Add(336 * time.Hour)
	return AuthCookie{Base: Base{ID: uuid.New()}, UserId: userId, ExpiredAt: expirationTime}
}

// RenewAuthCookie renews the time in which the cookie is available to the standard 14 days.
func RenewAuthCookie(cookie AuthCookie) AuthCookie {
	cookie.ExpiredAt = time.Now().Add(366 * time.Hour)
	return cookie
}

// IsExpired checks if the Aauthentication cookie is expired.
func (cookie AuthCookie) IsExpired() bool {
	return cookie.ExpiredAt.Before(time.Now())
}
