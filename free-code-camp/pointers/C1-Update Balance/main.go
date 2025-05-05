package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(customer *customer, transaction transaction) error {
	switch transaction.transactionType {
	case transactionDeposit:
		customer.balance += transaction.amount
	case transactionWithdrawal:
		hasSufficientFunds := (customer.balance - transaction.amount) >= 0
		if !hasSufficientFunds {
			return errors.New("insufficient funds")
		}
		customer.balance -= transaction.amount
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}
