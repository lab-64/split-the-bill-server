package model

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testdata
var (
	TestBill = BillModel{
		ID:    uuid.New(),
		Name:  "Test Bill",
		Owner: TestUser,
		Items: []ItemModel{TestItem1, TestItem2},
	}

	TestItem1 = ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 1",
		Price:        10,
		Contributors: []UserModel{TestUser},
	}

	TestItem2 = ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 2",
		Price:        18.5,
		Contributors: []UserModel{TestUser, TestUser2},
	}

	TestBill2 = BillModel{
		ID:    uuid.New(),
		Name:  "Test Bill 2",
		Owner: TestUser2,
		Items: []ItemModel{TestItem3},
	}

	TestItem3 = ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 3",
		Price:        30,
		Contributors: []UserModel{TestUser},
	}
)

func TestBillBalanceCalculation(t *testing.T) {
	balance := TestBill.CalculateBalance()
	assert.Equal(t, 2, len(balance))
	assert.Equal(t, 9.25, balance[TestUser.ID])
	assert.Equal(t, -9.25, balance[TestUser2.ID])
}
