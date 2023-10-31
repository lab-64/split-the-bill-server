package entity

import (
	"log"
	"split-the-bill-server/domain/model"
)

// User struct
type User struct {
	Base
	Username         string            `gorm:"unique;not null"`
	Groups           []*Group          `gorm:"many2many:group_members;"`
	GroupInvitations []GroupInvitation `gorm:"foreignKey:InviteeID"`
}

func MakeUser(user model.User) User {
	return User{Base: Base{ID: user.ID}, Username: user.Username}
}

// TODO convert groups
func (user *User) ToUser() model.User {
	var groupInvitations []model.GroupInvitation
	for _, groupInv := range user.GroupInvitations {
		j := groupInv.ToGroupInvitation()
		log.Println("Groups")
		log.Println(j.Group)
		groupInvitations = append(groupInvitations, groupInv.ToGroupInvitation())
	}

	return model.User{ID: user.ID, Username: user.Username, Groups: nil, PendingGroupInvitations: groupInvitations}
}

func ToUserSlice(users []User) []model.User {
	s := make([]model.User, len(users))
	for i, user := range users {
		s[i] = user.ToUser()
	}
	return s
}
