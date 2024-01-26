package model

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testdata
var (
	TestUser = UserModel{
		ID:    uuid.New(),
		Email: "test@mail.com",
	}
	TestUser2 = UserModel{
		ID:    uuid.New(),
		Email: "test2@mail.com",
	}

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

	TestGroup = GroupModel{
		ID:      uuid.New(),
		Name:    "Test Group",
		Owner:   TestUser,
		Members: []UserModel{TestUser, TestUser2},
		Bills:   []BillModel{TestBill},
	}
)

func TestBalanceCalculation(t *testing.T) {

	balance := TestGroup.CalculateBalance()
	assert.Equalf(t, 2, len(balance), "Balance should inclue 2 group members")
	assert.Equalf(t, 9.25, balance[TestUser.ID], "Balance for TestUser should be 9.25")
	assert.Equalf(t, -9.25, balance[TestUser2.ID], "Balance for TestUser2 should be -9.25")

}
