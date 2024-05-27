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

func CreateItem(id uuid.UUID, billID uuid.UUID, item dto.ItemInput) Item {
	// convert contributorIDs to simple UserModels
	var contributors []User
	if item.Contributors != nil {
		if item.Contributors.Add != nil {
			for _, contributorID := range *item.Contributors.Add {
				contributors = append(contributors, User{ID: contributorID})
			}
		}
	}

	return Item{
		ID:           id,
		Name:         *item.Name,
		Price:        *item.Price,
		BillID:       billID,
		Contributors: contributors,
	}
}

func (item *Item) UpdateItem(itemDTO dto.ItemInput) {
	// handle base fields
	if itemDTO.Name != nil {
		item.Name = *itemDTO.Name
	}
	if itemDTO.Price != nil {
		item.Price = *itemDTO.Price
	}
	// handle contributor changes
	if itemDTO.Contributors != nil {
		if itemDTO.Contributors.Add != nil {
			for _, contributorID := range *itemDTO.Contributors.Add {
				// check if contributor is already in the list
				_, userInLst := getUserByID(item.Contributors, contributorID)
				if userInLst == nil {
					item.Contributors = append(item.Contributors, User{ID: contributorID})
				}
			}
		}
		if itemDTO.Contributors.Remove != nil {
			for _, contributorID := range *itemDTO.Contributors.Remove {
				i, userInLst := getUserByID(item.Contributors, contributorID)
				// check if contributor is in the list
				if userInLst != nil {
					item.Contributors = append(item.Contributors[:i], item.Contributors[i+1:]...)
				}
			}
		}
		// update contributor uuids is not possible therefore no update needed
	}
}

// getUserByID returns the index and the user with the given id in the list of users
func getUserByID(users []User, id uuid.UUID) (int, *User) {
	for i, user := range users {
		if user.ID == id {
			return i, &user
		}
	}
	return -1, nil
}
