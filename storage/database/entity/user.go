package entity

import (
	"split-the-bill-server/domain/model"
)

type User struct {
	Base
	Email    string `gorm:"unique;not null"`
	Username string
	Groups   []*Group `gorm:"many2many:group_members;"`
}

func CreateUserEntity(user model.UserModel) User {
	return User{
		Base:     Base{ID: user.ID},
		Email:    user.Email,
		Username: user.Username,
	}
}

func ConvertToUserModel(user User) model.UserModel {
	return model.UserModel{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func ToUserModelSlice(users []User) []model.UserModel {
	s := make([]model.UserModel, len(users))
	for i, user := range users {
		s[i] = ConvertToUserModel(user)
	}
	return s
}
