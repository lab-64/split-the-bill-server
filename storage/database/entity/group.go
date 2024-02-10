package entity

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type Group struct {
	Base
	Name            string          `gorm:"not null"`
	OwnerUID        uuid.UUID       `gorm:"type:uuid"`
	Owner           User            `gorm:"foreignKey:OwnerUID"`
	Members         []*User         `gorm:"many2many:group_members"`
	Bills           []Bill          `gorm:"foreignKey:GroupID"` // has many bills
	GroupInvitation GroupInvitation `gorm:"foreignKey:GroupID"` // has one invitation
}

func CreateGroupEntity(group GroupModel) Group {
	// convert uuids to users
	var members []*User
	for _, member := range group.Members {
		members = append(members, &User{Base: Base{ID: member.ID}})
	}
	return Group{
		Base:            Base{ID: group.ID},
		OwnerUID:        group.Owner.ID,
		Name:            group.Name,
		Members:         members,
		GroupInvitation: GroupInvitation{Base: Base{ID: group.InvitationID}},
	}
}

func ConvertToGroupModel(group Group) GroupModel {
	members := make([]UserModel, len(group.Members))
	for i, member := range group.Members {
		members[i] = ConvertToUserModel(*member)
	}
	bills := make([]BillModel, len(group.Bills))
	for i, bill := range group.Bills {
		bills[i] = ConvertToBillModel(bill)
	}

	return GroupModel{
		ID:           group.ID,
		Name:         group.Name,
		Owner:        ConvertToUserModel(group.Owner),
		Members:      members,
		Bills:        bills,
		InvitationID: group.GroupInvitation.ID,
	}
}

func ToGroupModelSlice(groups []Group) []GroupModel {
	s := make([]GroupModel, len(groups))
	for i, group := range groups {
		s[i] = ConvertToGroupModel(group)
	}
	return s
}
