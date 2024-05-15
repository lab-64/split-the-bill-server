package entity

import (
	"github.com/google/uuid"
	"time"
)

type GroupTransaction struct {
	Base
	Date         time.Time     `gorm:"not null"`
	GroupID      uuid.UUID     `gorm:"type:uuid"` // group transaction belongs to a group
	Group        Group         `gorm:"foreignKey:GroupID"`
	Transactions []Transaction `gorm:"foreignKey:GroupTransactionID"`
}

type Transaction struct {
	Base
	DebtorID           uuid.UUID // transaction belongs to a debtor
	Debtor             User
	CreditorID         uuid.UUID // transaction belongs to a creditor
	Creditor           User
	Amount             float64
	GroupTransactionID uuid.UUID
}
