package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type Group struct {
	Base
	Owner   uuid.UUID
	User    User    `gorm:"foreignKey:Owner"`
	Name    string  `gorm:"not null"`
	Members []*User `gorm:"many2many:group_members;"`
}

// MakeGroup creates a database Group entity from a model.Group
func MakeGroup(group model.Group) Group {
	var members []*User
	for i := range group.Members {
		user := MakeUser(group.Members[i])
		members = append(members, &user)
	}
	return Group{Base: Base{ID: group.ID}, Owner: group.Owner.ID, Name: group.Name, Members: members}
}

// ToGroup creates a model.Group from a database Group entity
func (group *Group) ToGroup() model.Group {
	var members []model.User
	for i := range group.Members {
		members = append(members, group.Members[i].ToUser())
	}

	return model.Group{ID: group.ID, Owner: group.User.ToUser(), Name: group.Name, Members: members}
}
