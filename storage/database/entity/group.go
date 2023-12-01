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

func ToGroupEntity(group GroupModel) Group {
	// convert uuids to users
	var members []*User
	for _, member := range group.Members {
		members = append(members, &User{Base: Base{ID: member}})
	}
	// convert uuids to bills
	var bills []Bill
	for _, bill := range group.Bills {
		bills = append(bills, Bill{Base: Base{ID: bill}})
	}
	return Group{Base: Base{ID: group.ID}, OwnerUID: group.Owner, Name: group.Name, Members: members, Bills: bills}
}

func ToGroupModel(group *Group) GroupModel {
	// convert users to uuids
	var members []uuid.UUID
	for _, member := range group.Members {
		members = append(members, member.ID)
	}
	// convert bills to uuids
	var billIDs []uuid.UUID
	for _, bill := range group.Bills {
		billIDs = append(billIDs, bill.ID)
	}
	return GroupModel{ID: group.ID, Owner: group.OwnerUID, Name: group.Name, Members: members, Bills: billIDs}
}
