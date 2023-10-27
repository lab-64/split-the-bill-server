package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type Group struct {
	Base
	Owner   uuid.UUID
	User    User    `gorm:"foreignKey:Owner"`
	Name    string  `gorm:"not null"`
	Members []*User `gorm:"many2many:group_members;"`
}

// MakeGroup creates a database Group entity from a types.Group
func MakeGroup(group types.Group) Group {
	var members []*User
	for i := range group.Members {
		user := MakeUser(group.Members[i])
		members = append(members, &user)
	}
	return Group{Base: Base{ID: group.ID}, Owner: group.Owner.ID, Name: group.Name, Members: members}
}

// ToGroup creates a types.Group from a database Group entity
func (group *Group) ToGroup() types.Group {
<<<<<<< HEAD
	log.Println("user id ", group.User.ID)
=======
>>>>>>> dev
	var members []types.User
	for i := range group.Members {
		members = append(members, group.Members[i].ToUser())
	}

	return types.Group{ID: group.ID, Owner: group.User.ToUser(), Name: group.Name, Members: members}
}
