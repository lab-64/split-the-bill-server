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
	Items            []*Item           `gorm:"many2many:item_contributors"` // many to many
}

func ToUserEntity(user UserModel) User {
	return User{Base: Base{ID: user.ID}, Username: user.Username}
}

func ToUserModel(user User) UserModel {

	// convert groups
	var groups []GroupModel
	for _, group := range user.Groups {
		groups = append(groups, ToGroupModel(group))
	}

	// convert group invitations
	var groupInvitations []GroupInvitationModel
	for _, groupInv := range user.GroupInvitations {
		groupInvitations = append(groupInvitations, ToGroupInvitationModel(groupInv))
	}

	return UserModel{ID: user.ID, Username: user.Username, Groups: groups, PendingGroupInvitations: groupInvitations}
}

func ToUserModelSlice(users []User) []UserModel {
	s := make([]UserModel, len(users))
	for i, user := range users {
		s[i] = ToUserModel(user)
	}
	return s
}
