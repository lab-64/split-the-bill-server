package entity

import (
	"github.com/google/uuid"
	"time"
)

type Bill struct {
	Base
	OwnerID    uuid.UUID `gorm:"type:uuid"`
	Owner      User      `gorm:"foreignKey:OwnerID"` // belongs to user
	Name       string    `gorm:"not null"`
	Date       time.Time
	Items      []Item    `gorm:"foreignKey:BillID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // has many items
	GroupID    uuid.UUID `gorm:"type:uuid"`                                                       // group has many bills
	UnseenFrom []User    `gorm:"many2many:unseen_bills"`                                          // many to many unseen bills
}
