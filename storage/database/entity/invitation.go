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

func CreateGroupInvitationEntity(groupInvitation GroupInvitationModel) GroupInvitation {
	return GroupInvitation{
		Base:      Base{ID: groupInvitation.ID},
		Date:      groupInvitation.Date,
		GroupID:   groupInvitation.Group.ID,
		InviteeID: groupInvitation.Invitee.ID,
	}
}

func ConvertToGroupInvitationModel(inv GroupInvitation, isDetailed bool) GroupInvitationModel {
	group := GroupModel{}
	if isDetailed {
		group = ConvertToGroupModel(inv.For, false)
	}
	return GroupInvitationModel{
		ID:    inv.ID,
		Date:  inv.Date,
		Group: group,
	}
}
