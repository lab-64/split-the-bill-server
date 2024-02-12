package entity

import (
	"github.com/google/uuid"
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
