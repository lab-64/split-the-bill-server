package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

type GroupInvitation struct {
	Base
	Date    time.Time `gorm:"not null"`
	GroupID uuid.UUID `gorm:"not null"`
	For     Group     `gorm:"foreignKey:GroupID"`
}

// MakeGroupInvitation creates a database GroupInvitation entity from a types.GroupInvitation
func MakeGroupInvitation(groupInvitation types.GroupInvitation) GroupInvitation {
	group := MakeGroup(groupInvitation.For)
	return GroupInvitation{Base: Base{ID: groupInvitation.ID}, Date: groupInvitation.Date, GroupID: groupInvitation.For.ID, For: group}
}

// ToGroupInvitation creates a types.GroupInvitation from a database GroupInvitation entity
func (groupInvitation *GroupInvitation) ToGroupInvitation() types.GroupInvitation {
	group := groupInvitation.For.ToGroup()
	return types.GroupInvitation{ID: groupInvitation.ID, Date: groupInvitation.Date, For: group}
}
