package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Group struct {
	Base
	Name     string    `gorm:"not null"`
	OwnerUID uuid.UUID `gorm:"type:uuid"`
	Owner    User      `gorm:"foreignKey:OwnerUID"`
	Members  []*User   `gorm:"many2many:group_members"`
}

func ToGroupEntity(group GroupModel) Group {
	// convert uuids to users
	var members []*User
	for _, member := range group.Members {
		members = append(members, &User{Base: Base{ID: member}})
	}
	return Group{Base: Base{ID: group.ID}, OwnerUID: group.Owner, Name: group.Name, Members: members}
}

func ToGroupModel(group *Group) GroupModel {
	// convert users to uuids
	var members []uuid.UUID
	for _, member := range group.Members {
		members = append(members, member.ID)
	}
	return GroupModel{ID: group.ID, Owner: group.OwnerUID, Name: group.Name, Members: members}
}
