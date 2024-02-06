package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type GroupInvitation struct {
	Base
	GroupID uuid.UUID `gorm:"type:uuid"`
	Group   Group     `gorm:"constraint:OnDelete:CASCADE;"` // has one group, if group gets deleted, the invitation gets deleted too
}

func CreateGroupInvitationEntity(groupInvitation GroupInvitationModel) GroupInvitation {
	return GroupInvitation{
		Base:    Base{ID: groupInvitation.ID},
		GroupID: groupInvitation.Group.ID,
		Group:   CreateGroupEntity(groupInvitation.Group),
	}
}

func ConvertToGroupInvitationModel(inv GroupInvitation) GroupInvitationModel {

	return GroupInvitationModel{
		ID:    inv.ID,
		Group: ConvertToGroupModel(inv.Group),
	}
}
