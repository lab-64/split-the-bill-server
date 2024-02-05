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
	ID           uuid.UUID           `json:"id"`
	Name         string              `json:"name"`
	Price        float64             `json:"price"`
	BillID       uuid.UUID           `json:"billId"`
	Contributors []UserCoreOutputDTO `json:"contributors"`
}

func CreateItemModel(id uuid.UUID, item ItemInputDTO) ItemModel {
	// convert contributorIDs to simple UserModels
	contributors := make([]UserModel, len(item.Contributors))
	for i, contributorID := range item.Contributors {
		contributors[i] = UserModel{ID: contributorID}
	}
	return ItemModel{
		ID:           id,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}

}

func ConvertToItemDTO(item ItemModel) ItemOutputDTO {
	contributors := make([]UserCoreOutputDTO, len(item.Contributors))
	for i, cont := range item.Contributors {
		contributors[i] = ConvertToUserCoreDTO(&cont)
	}

	return ItemOutputDTO{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}
