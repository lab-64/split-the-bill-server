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

func CreateBill(id uuid.UUID, b dto.BillInput, date time.Time, unseenFrom []uuid.UUID) Bill {
	// convert each item
	var items []Item
	for _, item := range b.Items {
		items = append(items, CreateItem(uuid.New(), item))
	}
	return Bill{
		ID:               id,
		Owner:            User{ID: b.OwnerID},
		Name:             b.Name,
		Date:             date,
		GroupID:          b.GroupID,
		Items:            items,
		UnseenFromUserID: unseenFrom,
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

func (bill Bill) IsUnseen(userID uuid.UUID) bool {
	for _, id := range bill.UnseenFromUserID {
		if id == userID {
			return true
		}
	}
	return false
}
