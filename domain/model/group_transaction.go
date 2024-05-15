package model

import (
	"github.com/google/uuid"
	"math"
	"time"
)

type GroupTransaction struct {
	ID           uuid.UUID
	Date         time.Time
	GroupID      uuid.UUID
	GroupName    string
	Transactions []Transaction
}

type Transaction struct {
	ID       uuid.UUID
	Debtor   User
	Creditor User
	Amount   float64
}

type Amount struct {
	User   uuid.UUID
	Amount float64
}

func CreateTransaction(debtorID uuid.UUID, creditorID uuid.UUID, amount float64) Transaction {
	return Transaction{
		ID:       uuid.New(),
		Debtor:   User{ID: debtorID},
		Creditor: User{ID: creditorID},
		Amount:   amount,
	}
}

// getHighestValue returns the pair(user, amount) with the highest balance
func getHighestValue(balance map[uuid.UUID]float64) Amount {
	highest := Amount{
		Amount: 0.0,
	}
	for key, value := range balance {
		if value > highest.Amount {
			highest.User = key
			highest.Amount = value
		}
	}
	return highest
}

// getLowestValue returns the pair (user, amount) with the lowest balance
func getLowestValue(balance map[uuid.UUID]float64) Amount {
	lowest := Amount{
		Amount: 0.0,
	}
	for key, value := range balance {
		if value < lowest.Amount {
			lowest.User = key
			lowest.Amount = value
		}
	}
	return lowest
}

// splitBalance separates the balance into credit and debt
func splitBalance(balance map[uuid.UUID]float64) (map[uuid.UUID]float64, map[uuid.UUID]float64) {
	creditMap := make(map[uuid.UUID]float64)
	debtMap := make(map[uuid.UUID]float64)
	for k, v := range balance {
		if v > 0 {
			creditMap[k] = v
		}
		if v < 0 {
			debtMap[k] = v
		}
	}
	return creditMap, debtMap
}

// handleEqualAmount checks if there are equal amounts of credit and debt
func handleEqualAmount(creditMap *map[uuid.UUID]float64, debtMap *map[uuid.UUID]float64) []Transaction {
	var retSplit []Transaction
	for debtKey, debt := range *debtMap {
		for creditKey, credit := range *creditMap {
			if debt+credit == 0 { // if == 0 both have abs same amount
				retSplit = append(retSplit, CreateTransaction(debtKey, creditKey, credit))
				delete(*debtMap, debtKey)
				delete(*creditMap, creditKey)
				break
			}
		}
	}
	return retSplit
}

// ProduceTransactionsFromBalance produces transactions to clear a given balance.
// The returned transactions are the minimal amount of transactions needed to clear the balance.
// A transaction consists of debtors, their creditors and the amount to be repaid.
func ProduceTransactionsFromBalance(balance map[uuid.UUID]float64) []Transaction {
	var retSplit []Transaction
	remainingCredit, remainingDebt := splitBalance(balance)
	// check if there is no debt or credit
	if len(remainingDebt) == 0 || len(remainingCredit) == 0 {
		return retSplit
	}
	for range balance {
		equalSplit := handleEqualAmount(&remainingCredit, &remainingDebt)
		retSplit = append(retSplit, equalSplit...)
		// check if each debt or credit was handled
		if len(remainingDebt) == 0 || len(remainingCredit) == 0 {
			break
		}
		highestDebt := getLowestValue(remainingDebt)
		highestCred := getHighestValue(remainingCredit)
		clearAmount := math.Min(math.Abs(highestDebt.Amount), highestCred.Amount)
		retSplit = append(retSplit, CreateTransaction(highestDebt.User, highestCred.User, clearAmount))
		remainingDebt[highestDebt.User] = remainingDebt[highestDebt.User] + clearAmount
		if remainingDebt[highestDebt.User] == 0 {
			delete(remainingDebt, highestDebt.User)
		}
		remainingCredit[highestCred.User] = remainingCredit[highestCred.User] - clearAmount
		if remainingCredit[highestCred.User] == 0 {
			delete(remainingCredit, highestCred.User)
		}
		if len(remainingDebt) == 0 || len(remainingCredit) == 0 {
			break
		}
	}
	return retSplit
}
