package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type BillInputDTO struct {
	Owner   uuid.UUID      `json:"ownerID"`
	Name    string         `json:"name"`
	Date    time.Time      `json:"date"`
	GroupID uuid.UUID      `json:"groupID"`
	Items   []ItemInputDTO `json:"items"`
}

type BillDetailedOutputDTO struct {
	ID      uuid.UUID           `json:"id"`
	Name    string              `json:"name"`
	Date    time.Time           `json:"date"`
	Items   []ItemOutputDTO     `json:"items"`
	GroupID uuid.UUID           `json:"groupID"`
	Owner   UserPublicOutputDTO `json:"owner"`
}

func ToBillModel(b BillInputDTO) BillModel {
	// convert each item
	var items []ItemModel
	for _, item := range b.Items {
		items = append(items, ToItemModel(uuid.Nil, item))
	}
	return CreateBillModel(UserModel{ID: b.Owner}, b.Name, b.Date, b.GroupID, items)
}

func ToBillDetailedDTOs(bills []BillModel) []BillDetailedOutputDTO {
	billsDTO := make([]BillDetailedOutputDTO, len(bills))

	for i, bill := range bills {
		billsDTO[i] = ToBillDetailedDTO(bill)
	}
	return billsDTO
}

func ToBillDetailedDTO(bill BillModel) BillDetailedOutputDTO {
	itemsDTO := make([]ItemOutputDTO, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(item)
	}

	return BillDetailedOutputDTO{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Items:   itemsDTO,
		Owner:   ToUserPublicDTO(bill.Owner),
		GroupID: bill.GroupID,
	}
}
