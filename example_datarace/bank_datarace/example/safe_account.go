package example

import (
	"sync"
	"time"
)

type SafeAccount struct {
	ID      int
	Balance int
	mu      sync.Mutex
}

func (a *SafeAccount) Deposit(amount int) {
	time.Sleep(time.Microsecond)
	a.mu.Lock()
	defer a.mu.Unlock()
	time.Sleep(time.Microsecond)
	a.Balance += amount
}

func (a *SafeAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	time.Sleep(time.Microsecond)
	if a.Balance >= amount {
		time.Sleep(time.Microsecond)
		a.Balance -= amount
		return true
	}
	return false
}

func SafeTransfer(from, to *SafeAccount, amount int) bool {
	if from.Withdraw(amount) {
		time.Sleep(time.Microsecond)
		to.Deposit(amount)
		return true
	}
	return false
}

func (a *SafeAccount) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	time.Sleep(time.Microsecond)
	return a.Balance
}
