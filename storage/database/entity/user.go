package entity

import (
	"split-the-bill-server/domain/model"
)

type User struct {
	Base
	Email            string            `gorm:"unique;not null"`
	Groups           []*Group          `gorm:"many2many:group_members;"`
	GroupInvitations []GroupInvitation `gorm:"foreignKey:InviteeID"`
}

func CreateUserEntity(user model.UserModel) User {
	return User{
		Base:  Base{ID: user.ID},
		Email: user.Email,
	}
}

func ConvertToUserModel(user User, isDetailed bool) model.UserModel {

	groupInvitations := make([]model.GroupInvitationModel, len(user.GroupInvitations))
	groupModels := make([]model.GroupModel, len(user.Groups))

	if isDetailed {
		for i, inv := range user.GroupInvitations {
			groupInvitations[i] = ConvertToGroupInvitationModel(inv, false)
		}

		for i, group := range user.Groups {
			groupModels[i] = ConvertToGroupModel(*group, false)
		}
	}

	return model.UserModel{
		ID:                      user.ID,
		Email:                   user.Email,
		PendingGroupInvitations: groupInvitations,
		Groups:                  groupModels,
	}
}

func ToUserModelSlice(users []User) []model.UserModel {
	s := make([]model.UserModel, len(users))
	for i, user := range users {
		s[i] = ConvertToUserModel(user, true)
	}
	return s
}
