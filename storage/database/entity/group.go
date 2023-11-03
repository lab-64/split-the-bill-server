package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type Group struct {
	Base
	Name     string    `gorm:"not null"`
	OwnerUID uuid.UUID `gorm:"type:uuid"`
	Owner    User      `gorm:"foreignKey:OwnerUID"`
	Members  []*User   `gorm:"many2many:group_members"`
}

// MakeGroup creates a database Group entity from a types.Group
func MakeGroup(group types.Group) Group {
	// convert uuids to users
	var members []*User
	for _, member := range group.Members {
		members = append(members, &User{Base: Base{ID: member}})
	}
	return Group{Base: Base{ID: group.ID}, OwnerUID: group.Owner, Name: group.Name, Members: members}
}

// ToGroup creates a types.Group from a database Group entity
func (group *Group) ToGroup() types.Group {
	// convert users to uuids
	var members []uuid.UUID
	for _, member := range group.Members {
		members = append(members, member.ID)
	}
	return types.Group{ID: group.ID, Owner: group.OwnerUID, Name: group.Name, Members: members}
}
