package model

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
	"time"
)

type Bill struct {
	ID               uuid.UUID
	Owner            User
	Name             string
	Date             time.Time
	GroupID          uuid.UUID
	Items            []Item
	Balance          map[uuid.UUID]float64
	UnseenFromUserID []uuid.UUID
}

func CreateBill(id uuid.UUID, ownerID uuid.UUID, name string, date time.Time, groupID uuid.UUID, items []Item, unseenFrom []uuid.UUID) Bill {
	return Bill{
		ID:               id,
		Owner:            User{ID: ownerID},
		Name:             name,
		Date:             date,
		GroupID:          groupID,
		Items:            items,
		UnseenFromUserID: unseenFrom,
	}
}

func (bill *Bill) UpdateBill(billDTO dto.BillUpdate) {
	// handle base fields
	if billDTO.Name != bill.Name {
		bill.Name = billDTO.Name
	}
	if billDTO.Date != bill.Date {
		bill.Date = billDTO.Date
	}
	// handle item changes
	var newItemLst []Item
	// handle item addition
	addedItems := getItemDTOsByID(billDTO.Items, uuid.Nil)
	for _, addedItem := range addedItems {
		newItemLst = append(newItemLst, CreateItem(uuid.New(), bill.ID, addedItem))
	}
	// handle item update
	for _, item := range bill.Items {
		for _, itemDTO := range billDTO.Items {
			// update item with the matching id
			if item.ID == itemDTO.ID {
				item.UpdateItem(itemDTO)
				break
			}
		}
		// items with no matching IDs are removed from the user and therefore are not added to the new item list
	}
	bill.Items = newItemLst
}

func (bill *Bill) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	for _, item := range bill.Items {
		ppp := item.Price / float64(len(item.Contributors))
		for _, contributor := range item.Contributors {
			balance[contributor.ID] -= ppp
		}
		balance[bill.Owner.ID] += item.Price
	}
	return balance
}

func (bill *Bill) IsUnseen(userID uuid.UUID) bool {
	for _, id := range bill.UnseenFromUserID {
		if id == userID {
			return true
		}
	}
	return false
}

// getItemDTOsByID returns all items with the given id
func getItemDTOsByID(items []dto.ItemInput, id uuid.UUID) []dto.ItemInput {
	var output []dto.ItemInput
	for _, item := range items {
		if item.ID == id {
			output = append(output, item)
		}
	}
	return output
}
