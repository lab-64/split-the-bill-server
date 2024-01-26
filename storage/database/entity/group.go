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
	Bills    []Bill    `gorm:"foreignKey:GroupID"` // has many bills
}

func CreateGroupEntity(group GroupModel) Group {
	// convert uuids to users
	var members []*User
	for _, member := range group.Members {
		members = append(members, &User{Base: Base{ID: member.ID}})
	}
	return Group{Base: Base{ID: group.ID}, OwnerUID: group.Owner.ID, Name: group.Name, Members: members}
}

func ConvertToGroupModel(group Group, isDetailed bool) GroupModel {
	members := make([]UserModel, len(group.Members))
	bills := make([]BillModel, len(group.Bills))
	owner := UserModel{}

	if isDetailed {

		for i, member := range group.Members {
			members[i] = ConvertToUserModel(*member, false)
		}

		for i, bill := range group.Bills {
			bills[i] = ConvertToBillModel(bill, false)
		}
		owner = ConvertToUserModel(group.Owner, false)
	}

	return GroupModel{
		ID:      group.ID,
		Name:    group.Name,
		Owner:   owner,
		Members: members,
		Bills:   bills,
	}
}

func ToGroupModelSlice(groups []Group) []GroupModel {
	s := make([]GroupModel, len(groups))
	for i, group := range groups {
		s[i] = ConvertToGroupModel(group, true)
	}
	return s
}
