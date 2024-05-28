package model

import (
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/presentation/dto"
)

type Item struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	BillID       uuid.UUID
	Contributors []User
}

func CreateItem(id uuid.UUID, billID uuid.UUID, item dto.ItemInput) Item {
	// convert contributorIDs to simple UserModels
	contributors := make([]User, len(item.Contributors))
	for i, contributorID := range item.Contributors {
		contributors[i] = User{ID: contributorID}
	}
	return Item{
		ID:           id,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       billID,
		Contributors: contributors,
	}
}

func (item *Item) UpdateItem(itemDTO dto.ItemInput) {
	log.Println("Update Item with ID: ", item.ID)
	// handle base fields
	if itemDTO.Name != item.Name {
		item.Name = itemDTO.Name
	}
	if itemDTO.Price != item.Price {
		item.Price = itemDTO.Price
	}
	// handle contributor changes
}
