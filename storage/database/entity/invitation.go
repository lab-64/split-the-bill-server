package entity

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"time"
)

type GroupInvitation struct {
	Base
	Date      time.Time `gorm:"not null"`
	GroupID   uuid.UUID `gorm:"not null"`
	For       Group     `gorm:"foreignKey:GroupID"`
	InviteeID uuid.UUID
}

// MakeGroupInvitation creates a database GroupInvitation entity from a model.GroupInvitation.
func MakeGroupInvitation(groupInvitation model.GroupInvitation) GroupInvitation {
	return GroupInvitation{Base: Base{ID: groupInvitation.ID}, Date: groupInvitation.Date, GroupID: groupInvitation.Group.ID, InviteeID: groupInvitation.Invitee.ID}
}

// ToGroupInvitation creates a model.GroupInvitation from a database GroupInvitation entity.
func (groupInvitation *GroupInvitation) ToGroupInvitation() model.GroupInvitation {
	group := groupInvitation.For.ToGroup()

	return model.GroupInvitation{ID: groupInvitation.ID, Date: groupInvitation.Date, Group: group}
}
