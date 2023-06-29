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
	Username string `json:"username"`
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