package model

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testdata
var (
	TestUser = User{
		ID:    uuid.New(),
		Email: "test@mail.com",
	}
	TestUser2 = User{
		ID:    uuid.New(),
		Email: "test2@mail.com",
	}
	TestUser3 = User{
		ID:    uuid.New(),
		Email: "test3@mail.com",
	}

	TestGroup = Group{
		ID:      uuid.New(),
		Name:    "Test Group",
		Owner:   TestUser,
		Members: []User{TestUser, TestUser2, TestUser3},
		Bills:   []Bill{TestBill, TestBill2},
	}

	TestGroup2 = Group{
		ID:      uuid.New(),
		Name:    "Test Group2",
		Owner:   TestUser,
		Members: []User{TestUser, TestUser2, TestUser3},
		Bills:   []Bill{TestBill, TestBill2, TestBill3},
	}
)

func TestGroup_CalculateBalance(t *testing.T) {
	balance := TestGroup.CalculateBalance()
	assert.Equal(t, 3, len(balance))
	assert.Equal(t, -20.75, balance[TestUser.ID])
	assert.Equal(t, 20.75, balance[TestUser2.ID])
	assert.Equal(t, 0.0, balance[TestUser3.ID])
}

func TestGroup_CalculateBalance_Empty_Contribution(t *testing.T) {
	// group with the same bills as TestGroup and an additional bill with no contributors should return the same balance as TestGroup
	balance := TestGroup2.CalculateBalance()
	assert.Equal(t, 3, len(balance))
	assert.Equal(t, -20.75, balance[TestUser.ID])
	assert.Equal(t, 20.75, balance[TestUser2.ID])
	assert.Equal(t, 0.0, balance[TestUser3.ID])
}

func TestGroup_IsMember(t *testing.T) {
	assert.True(t, TestGroup.IsMember(TestUser.ID))
	assert.True(t, TestGroup.IsMember(TestUser2.ID))
	assert.True(t, TestGroup.IsMember(TestUser3.ID))
	assert.False(t, TestGroup.IsMember(uuid.New()))
}

func TestGroup_SortBillsByDate(t *testing.T) {
	TestGroup.SortBillsByDate()
	assert.Equal(t, TestBill2.ID, TestGroup.Bills[0].ID)
	assert.Equal(t, TestBill.ID, TestGroup.Bills[1].ID)
}
