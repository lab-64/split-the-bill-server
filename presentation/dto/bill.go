package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
	"time"
)

type BaseBill struct {
	OwnerID uuid.UUID `json:"ownerID"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	GroupID uuid.UUID `json:"groupID"`
	Items   []Item    `json:"items"`
}

type Bill struct {
	BaseBill
	ID uuid.UUID `json:"id"`
}

func CreateBill(id uuid.UUID, owner uuid.UUID, name string, data time.Time, groupID uuid.UUID, items []Item) Bill {
	return Bill{
		ID: id,
		BaseBill: BaseBill{
			OwnerID: owner,
			Name:    name,
			Date:    data,
			GroupID: groupID,
			Items:   items,
		},
	}
}

type BillDetailedOutputDTO struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Items   []Item    `json:"items"`
	GroupID uuid.UUID `json:"groupID"`
	OwnerID uuid.UUID `json:"ownerID"`
}

func ToBillModel(id uuid.UUID, bill BaseBill, items []ItemModel) BillModel {
	return CreateBillModel(id, bill.OwnerID, bill.Name, bill.Date, bill.GroupID, items)
}

func ToBillDTO(bill BillModel) Bill {
	// convert each item
	var items []Item
	for _, item := range bill.Items {
		items = append(items, ToItemDTO(item))
	}
	return CreateBill(bill.ID, bill.OwnerID, bill.Name, bill.Date, bill.GroupID, items)
}

func ToBillDetailedDTOs(bills []BillModel) []BillDetailedOutputDTO {
	billsDTO := make([]BillDetailedOutputDTO, len(bills))

	for i, bill := range bills {
		billsDTO[i] = ToBillDetailedDTO(bill)
	}
	return billsDTO
}

func ToBillDetailedDTO(bill BillModel) BillDetailedOutputDTO {
	itemsDTO := make([]Item, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(item)
	}

	return BillDetailedOutputDTO{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Items:   itemsDTO,
		OwnerID: bill.OwnerID,
		GroupID: bill.GroupID,
	}
}
