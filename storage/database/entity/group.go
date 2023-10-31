package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Group struct {
	Base
	Owner   uuid.UUID
	User    User    `gorm:"foreignKey:Owner"`
	Name    string  `gorm:"not null"`
	Members []*User `gorm:"many2many:group_members;"`
}

func ToGroupEntity(group GroupModel) Group {
	var members []*User
	for _, member := range group.Members {
		user := ToUserEntity(member)
		members = append(members, &user)
	}

	return Group{
		Base:    Base{ID: group.ID},
		Owner:   group.Owner.ID,
		Name:    group.Name,
		Members: members,
	}
}

func ToGroupModel(group *Group) GroupModel {

	var members []UserModel
	for _, member := range group.Members {
		members = append(members, ToUserModel(*member))
	}

	return GroupModel{
		ID:      group.ID,
		Owner:   ToUserModel(group.User),
		Name:    group.Name,
		Members: members,
	}
}
