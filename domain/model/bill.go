package model

import (
	"github.com/google/uuid"
	"time"
)

type Bill struct {
	ID               uuid.UUID
	UpdatedAt        time.Time
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

func (bill *Bill) CalculateBalance() map[uuid.UUID]float64 {
	balance := make(map[uuid.UUID]float64)
	for _, item := range bill.Items {
		// zero check
		if len(item.Contributors) == 0 {
			continue
		}
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
