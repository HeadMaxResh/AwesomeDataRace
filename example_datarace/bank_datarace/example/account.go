package example

import "time"

type Account struct {
	ID      int
	Balance int
}

func (a *Account) Deposit(amount int) {
	balance := a.Balance
	time.Sleep(time.Microsecond)
	a.Balance = balance + amount
}

func (a *Account) Withdraw(amount int) bool {
	balance := a.Balance
	if balance >= amount {
		time.Sleep(time.Microsecond)
		a.Balance = balance - amount
		return true
	}
	return false
}

func Transfer(from, to *Account, amount int) bool {
	time.Sleep(time.Microsecond)
	if from.Withdraw(amount) {
		time.Sleep(time.Microsecond)
		to.Deposit(amount)
		return true
	}
	return false
}
