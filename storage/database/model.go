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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

// Create a database user. Password gets hashed.
func MakeUser(user types.User) User {

	// Hash Password
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	return User{Base: Base{ID: user.ID}, Email: user.Email, Password: password}
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
