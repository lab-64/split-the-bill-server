package model

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TestBalance = map[uuid.UUID]float64{
		uuid.New(): 0,
		uuid.New(): 20,
		uuid.New(): 10,
		uuid.New(): -10,
		uuid.New(): -13,
		uuid.New(): -7,
		uuid.New(): 2,
		uuid.New(): 8,
		uuid.New(): -10,
	}

	TestBalance2 = map[uuid.UUID]float64{
		uuid.New(): 20,
		uuid.New(): -10,
		uuid.New(): -13,
		uuid.New(): -7,
		uuid.New(): 2,
		uuid.New(): 8,
	}

	TestBalance3 = map[uuid.UUID]float64{
		uuid.New(): 13.2,
		uuid.New(): 10.8,
		uuid.New(): -114.1,
		uuid.New(): 90.1,
		uuid.New(): -5.2,
		uuid.New(): 13.21,
		uuid.New(): -65.21,
		uuid.New(): 16.1,
		uuid.New(): 54.21,
		uuid.New(): 31.12,
		uuid.New(): -45.1,
		uuid.New(): -31.25,
		uuid.New(): -6.56,
		uuid.New(): -4.12,
		uuid.New(): -67.21,
		uuid.New(): 110.01,
	}

	TestEmptyBalance = map[uuid.UUID]float64{
		uuid.New(): 0,
		uuid.New(): 0,
		uuid.New(): 0,
	}
)

func TestSplitGroupBalance(t *testing.T) {

	tests := []struct {
		name    string
		balance map[uuid.UUID]float64
	}{
		{
			name:    "Success",
			balance: TestBalance,
		},
		{
			name:    "Success Empty Balance",
			balance: TestEmptyBalance,
		},
		{
			name:    "Success 2",
			balance: TestBalance2,
		},
		{
			name:    "Success 3 - Big Group",
			balance: TestBalance3,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			splits := ProduceTransactionsFromBalance(testcase.balance)
			inputBalance := testcase.balance
			for _, splitEntry := range splits {
				inputBalance[splitEntry.Debtor.ID] += splitEntry.Amount
				inputBalance[splitEntry.Creditor.ID] -= splitEntry.Amount
				// check if no split contains a value of 0
				assert.NotEqualf(t, 0.0, splitEntry.Amount, "Split must not contain a value of 0")
			}
			// check if balance of all users is 0 => all debts are paid
			for _, balance := range TestBalance {
				assert.Equalf(t, 0.0, balance, "Balance must be 0")
			}
		})
	}
}
