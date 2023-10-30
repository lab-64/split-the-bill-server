package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"time"
)

type GroupInvitation struct {
	Base
	Date      time.Time `gorm:"not null"`
	GroupID   uuid.UUID `gorm:"not null"`
	For       Group     `gorm:"foreignKey:GroupID"`
	InviteeID uuid.UUID `gorm:"not null"`
	Invitee   User      `gorm:"foreignKey:InviteeID"`
}

// MakeGroupInvitation creates a database GroupInvitation entity from a types.GroupInvitation.
func MakeGroupInvitation(groupInvitation types.GroupInvitation) GroupInvitation {
	return GroupInvitation{Base: Base{ID: groupInvitation.ID}, Date: groupInvitation.Date, GroupID: groupInvitation.Group.ID, InviteeID: groupInvitation.Invitee.ID}
}

// ToGroupInvitation creates a types.GroupInvitation from a database GroupInvitation entity.
func (groupInvitation *GroupInvitation) ToGroupInvitation() types.GroupInvitation {
	group := groupInvitation.For.ToGroup()
	invitee := groupInvitation.Invitee.ToUser()

	return types.GroupInvitation{ID: groupInvitation.ID, Date: groupInvitation.Date, Group: group, Invitee: invitee}
}
