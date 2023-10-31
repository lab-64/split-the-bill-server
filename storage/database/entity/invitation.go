package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type GroupInvitation struct {
	Base
	Date      time.Time `gorm:"not null"`
	GroupID   uuid.UUID `gorm:"not null"`
	For       Group     `gorm:"foreignKey:GroupID"`
	InviteeID uuid.UUID
}

func ToGroupInvitationModel(groupInvitation GroupInvitation) GroupInvitationModel {
	return GroupInvitationModel{
		ID:    groupInvitation.ID,
		Date:  groupInvitation.Date,
		Group: ToGroupModel(&groupInvitation.For),
	}
}

func ToGroupInvitationEntity(groupInvitation GroupInvitationModel) GroupInvitation {
	return GroupInvitation{
		Base:      Base{ID: groupInvitation.ID},
		Date:      groupInvitation.Date,
		GroupID:   groupInvitation.Group.ID,
		InviteeID: groupInvitation.Invitee.ID,
	}
}
