package entity

import "github.com/google/uuid"

type Credentials struct {
	UserID uuid.UUID `gorm:"type:uuid; column:user_foreign_key;not null"`
	Hash   []byte    `gorm:"type:bytea;not null"`
}
