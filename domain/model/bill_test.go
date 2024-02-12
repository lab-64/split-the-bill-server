package model

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testdata
var (
	TestBill = Bill{
		ID:    uuid.New(),
		Name:  "Test Bill",
		Owner: TestUser,
		Items: []Item{TestItem1, TestItem2},
	}

	TestItem1 = Item{
		ID:           uuid.New(),
		Name:         "Test Item 1",
		Price:        10,
		Contributors: []User{TestUser},
	}

	TestItem2 = Item{
		ID:           uuid.New(),
		Name:         "Test Item 2",
		Price:        18.5,
		Contributors: []User{TestUser, TestUser2},
	}

	TestBill2 = Bill{
		ID:    uuid.New(),
		Name:  "Test Bill 2",
		Owner: TestUser2,
		Items: []Item{TestItem3},
	}

	TestItem3 = Item{
		ID:           uuid.New(),
		Name:         "Test Item 3",
		Price:        30,
		Contributors: []User{TestUser},
	}
)

func TestBillBalanceCalculation(t *testing.T) {
	balance := TestBill.CalculateBalance()
	assert.Equal(t, 2, len(balance))
	assert.Equal(t, 9.25, balance[TestUser.ID])
	assert.Equal(t, -9.25, balance[TestUser2.ID])
}
