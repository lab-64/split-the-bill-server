package entity

import (
	"log"
	"split-the-bill-server/types"
)

// User struct
type User struct {
	Base
	Username         string            `gorm:"unique;not null"`
	Groups           []*Group          `gorm:"many2many:group_members;"`
	GroupInvitations []GroupInvitation `gorm:"foreignKey:InviteeID"`
}

func MakeUser(user types.User) User {
	return User{Base: Base{ID: user.ID}, Username: user.Username}
}

// TODO convert groups
func (user *User) ToUser() types.User {
	var groupInvitations []types.GroupInvitation
	for _, groupInv := range user.GroupInvitations {
		j := groupInv.ToGroupInvitation()
		log.Println("Groups")
		log.Println(j.Group)
		groupInvitations = append(groupInvitations, groupInv.ToGroupInvitation())
	}

	return types.User{ID: user.ID, Username: user.Username, Groups: nil, PendingGroupInvitations: groupInvitations}
}

func ToUserSlice(users []User) []types.User {
	s := make([]types.User, len(users))
	for i, user := range users {
		s[i] = user.ToUser()
	}
	return s
}
