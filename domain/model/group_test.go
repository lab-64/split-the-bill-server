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
	TestUser3 = UserModel{
		ID:    uuid.New(),
		Email: "test3@mail.com",
	}

	TestGroup = GroupModel{
		ID:      uuid.New(),
		Name:    "Test Group",
		Owner:   TestUser,
		Members: []UserModel{TestUser, TestUser2, TestUser3},
		Bills:   []BillModel{TestBill, TestBill2},
	}
)

func TestGroupBalanceCalculation(t *testing.T) {

	balance := TestGroup.CalculateBalance()
	assert.Equal(t, 3, len(balance))
	assert.Equal(t, -20.75, balance[TestUser.ID])
	assert.Equal(t, 20.75, balance[TestUser2.ID])
	assert.Equal(t, 0.0, balance[TestUser3.ID])
}
