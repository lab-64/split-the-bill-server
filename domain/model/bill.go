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
	if billDTO.Name != nil {
		bill.Name = *billDTO.Name
	}
	if billDTO.Date != nil {
		bill.Date = *billDTO.Date
	}
	// handle item changes
	if billDTO.Items != nil {
		if billDTO.Items.Add != nil {
			for _, itemDTO := range *billDTO.Items.Add {
				item := CreateItem(uuid.New(), bill.ID, itemDTO)
				bill.Items = append(bill.Items, item)
			}
		}
		if billDTO.Items.Remove != nil {
			for _, removableItem := range *billDTO.Items.Remove {
				index, item := getItemByID(bill.Items, removableItem.ID)
				if item != nil {
					bill.Items = append(bill.Items[:index], bill.Items[index+1:]...)
				}
			}
		}
		if billDTO.Items.Update != nil {
			for _, updatedItem := range *billDTO.Items.Update {
				index, item := getItemByID(bill.Items, updatedItem.ID)
				if item != nil {
					item.UpdateItem(updatedItem)
					bill.Items[index] = *item
				}
			}
		}
	}
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

// getItemByID returns the item with the given id and its position from the item slice
func getItemByID(items []Item, id uuid.UUID) (int, *Item) {
	for i, item := range items {
		if item.ID == id {
			return i, &item
		}
	}
	return -1, nil
}
