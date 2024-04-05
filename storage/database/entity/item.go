package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	Base
	Name         string    `gorm:"not null"`
	Price        float64   `gorm:"not null"`
	BillID       uuid.UUID `gorm:"type:uuid"`                                                                 // belongs to bill
	Contributors []*User   `gorm:"many2many:item_contributors; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // many to many
}
