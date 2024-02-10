package dto

import (
	"github.com/google/uuid"
	"time"
)

type BillInput struct {
	OwnerID uuid.UUID   `json:"ownerID"`
	Name    string      `json:"name"`
	Date    time.Time   `json:"date"`
	GroupID uuid.UUID   `json:"groupID"`
	Items   []ItemInput `json:"items"`
}

type BillDetailedOutput struct {
	ID      uuid.UUID             `json:"id"`
	Name    string                `json:"name"`
	Date    time.Time             `json:"date"`
	Items   []ItemOutput          `json:"items"`
	GroupID uuid.UUID             `json:"groupID"`
	Owner   UserCoreOutput        `json:"owner"`
	Balance map[uuid.UUID]float64 `json:"balance,omitempty"` // include balance only if balance is set
}
