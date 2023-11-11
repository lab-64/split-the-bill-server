package entity

import "github.com/google/uuid"

type Credentials struct {
	UserID uuid.UUID `gorm:"type:uuid; not null"`
	User   User
	Hash   []byte `gorm:"type:bytea;not null"`
}
