package entity

import (
	"split-the-bill-server/types"
)

// User struct
type User struct {
	Base
	Username string   `gorm:"unique;not null"`
	Groups   []*Group `gorm:"many2many:group_members;"`
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
