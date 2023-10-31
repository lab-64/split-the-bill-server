package entity

import (
	. "split-the-bill-server/domain/model"
)

// User struct
type User struct {
	Base
	Username         string            `gorm:"unique;not null"`
	Groups           []*Group          `gorm:"many2many:group_members;"`
	GroupInvitations []GroupInvitation `gorm:"foreignKey:InviteeID"`
}

func ToUserEntity(user UserModel) User {
	return User{Base: Base{ID: user.ID}, Username: user.Username}
}

// ToUserModel TODO convert groups
func ToUserModel(user User) UserModel {
	var groupInvitations []GroupInvitationModel
	for _, groupInv := range user.GroupInvitations {
		groupInvitations = append(groupInvitations, ToGroupInvitationModel(groupInv))
	}

	return UserModel{ID: user.ID, Username: user.Username, Groups: nil, PendingGroupInvitations: groupInvitations}
}

func ToUserModelSlice(users []User) []UserModel {
	s := make([]UserModel, len(users))
	for i, user := range users {
		s[i] = ToUserModel(user)
	}
	return s
}
