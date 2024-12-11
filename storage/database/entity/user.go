package entity

import "github.com/google/uuid"

type User struct {
	Base
	PrivateAuth    uuid.UUID `gorm:"type:uuid"` // used for account less registration // TODO: should be unique but then we need to generate a uuid as default
	Email          string    `gorm:"unique;not null"`
	Username       string    `gorm:"not null"`
	ProfileImgPath string
	Groups         []*Group `gorm:"many2many:group_members; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
