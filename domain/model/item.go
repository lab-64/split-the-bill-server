package model

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type Item struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	BillID       uuid.UUID
	Contributors []User
}

func CreateItem(id uuid.UUID, item dto.ItemInput) Item {
	// convert contributorIDs to simple UserModels
	contributors := make([]User, len(item.Contributors))
	for i, contributorID := range item.Contributors {
		contributors[i] = User{ID: contributorID}
	}
	return Item{
		ID:           id,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}
