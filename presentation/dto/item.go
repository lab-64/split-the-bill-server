package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type ItemInputDTO struct {
	Name         string      `json:"name"`
	Price        float64     `json:"price"`
	BillID       uuid.UUID   `json:"billId"`
	Contributors []uuid.UUID `json:"contributorIDs"`
}

type ItemOutputDTO struct {
	ID           uuid.UUID             `json:"id"`
	Name         string                `json:"name"`
	Price        float64               `json:"price"`
	BillID       uuid.UUID             `json:"billId"`
	Contributors []UserPublicOutputDTO `json:"contributors"`
}

func ToItemModel(id uuid.UUID, i ItemInputDTO) ItemModel {
	// convert contributorIDs to simple UserModels
	contributors := make([]UserModel, len(i.Contributors))
	for i, contributorID := range i.Contributors {
		contributors[i] = UserModel{ID: contributorID}
	}
	return CreateItemModel(id, i.Name, i.Price, contributors, i.BillID)
}

// ToItemDTO converts an ItemModel to an ItemOutputDTO
func ToItemDTO(item ItemModel) ItemOutputDTO {
	contributors := make([]UserPublicOutputDTO, len(item.Contributors))
	for i, cont := range item.Contributors {
		contributors[i] = ToUserPublicDTO(cont)
	}

	return ItemOutputDTO{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}
