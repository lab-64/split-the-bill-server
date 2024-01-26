package model

import (
	"github.com/google/uuid"
	"time"
)

type BillModel struct {
	ID      uuid.UUID
	Owner   UserModel
	Name    string
	Date    time.Time
	GroupID uuid.UUID
	Items   []ItemModel
}
