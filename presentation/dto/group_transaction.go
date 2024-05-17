package dto

import (
	"github.com/google/uuid"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type GroupTransactionOutput struct {
	ID           uuid.UUID           `json:"id"`
	Date         time.Time           `json:"date"`
	GroupID      uuid.UUID           `json:"groupId"`
	GroupName    string              `json:"groupName"`
	Transactions []TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	Debtor   UserCoreOutput `json:"debtor"`
	Creditor UserCoreOutput `json:"creditor"`
	Amount   float64        `json:"amount"`
}
