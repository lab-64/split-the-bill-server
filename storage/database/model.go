package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/types"
	"time"
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
	Username string `gorm:"unique;not null"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// AuthCookie struct
type AuthCookie struct {
	Base
	UserID      uuid.UUID `gorm:"type:uuid; column:user_foreign_key;not null"`
	ValidBefore time.Time
}

// Credentials struct
type Credentials struct {
	UserID uuid.UUID `gorm:"type:uuid; column:user_foreign_key;not null"`
	Hash   []byte    `gorm:"type:bytea;not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

func MakeUser(user types.User) User {
	return User{Base: Base{ID: user.ID}, Username: user.Username}
}

func (user *User) ToUser() types.User {
	return types.User{ID: user.ID, Username: user.Username}
}

func ToUserSlice(users []User) []types.User {
	s := make([]types.User, len(users))
	for i, user := range users {
		s[i] = user.ToUser()
	}
	return s
}

func MakeAuthCooke(authCookie types.AuthenticationCookie) AuthCookie {
	return AuthCookie{Base: Base{ID: authCookie.Token}, UserID: authCookie.UserID, ValidBefore: authCookie.ValidBefore}
}

func (authCookie *AuthCookie) ToAuthCookie() types.AuthenticationCookie {
	return types.AuthenticationCookie{UserID: authCookie.UserID, Token: authCookie.ID, ValidBefore: authCookie.ValidBefore}
}
