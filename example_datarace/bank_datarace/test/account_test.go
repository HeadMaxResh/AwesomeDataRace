package test

import (
	"awesomeDataRace/example_datarace/bank_datarace/example"
	"sync"
	"testing"
)

func TestAccount_WithdrawAndDepositRace(t *testing.T) {
	acc := &example.Account{ID: 1, Balance: 1000}
	var wg sync.WaitGroup

	// 100 пополнений по 10 и 100 снятий по 10 — итог должен быть 1000
	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			acc.Deposit(10)
		}()
		go func() {
			defer wg.Done()
			acc.Withdraw(10)
		}()
	}

	wg.Wait()

	if acc.Balance != 1000 {
		t.Errorf("Race detected: expected balance 1000, got %d", acc.Balance)
	}
}

func TestAccount_MultipleWithdrawals(t *testing.T) {
	acc := &example.Account{ID: 1, Balance: 100}
	var wg sync.WaitGroup
	failures := 0
	var mu sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !acc.Withdraw(10) {
				mu.Lock()
				failures++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	totalWithdrawn := 100 - acc.Balance
	if totalWithdrawn > 100 {
		t.Errorf("Overdrawn: total withdrawn = %d, balance = %d", totalWithdrawn, acc.Balance)
	} else {
		t.Logf("Withdrawn: %d, Failed ops: %d, Final balance: %d", totalWithdrawn, failures, acc.Balance)
	}
}

func TestAccount_ConcurrentDeposits(t *testing.T) {
	acc := &example.Account{ID: 1, Balance: 0}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Deposit(1)
		}()
	}
	wg.Wait()

	if acc.Balance != 1000 {
		t.Errorf("Incorrect final balance: expected 1000, got %d", acc.Balance)
	}
}

func TestAccount_TransferRace(t *testing.T) {
	acc1 := &example.Account{ID: 1, Balance: 1000}
	acc2 := &example.Account{ID: 2, Balance: 1000}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			example.Transfer(acc1, acc2, 10)
		}()
		go func() {
			defer wg.Done()
			example.Transfer(acc2, acc1, 10)
		}()
	}

	wg.Wait()

	total := acc1.Balance + acc2.Balance
	if total != 2000 {
		t.Errorf("Inconsistent total: expected 2000, got %d", total)
	} else {
		t.Logf("Acc1: %d, Acc2: %d, Total: %d", acc1.Balance, acc2.Balance, total)
	}
}

func TestAccount_DataCorruption(t *testing.T) {
	acc := &example.Account{ID: 1, Balance: 100}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			acc.Deposit(1)
		}()
		go func() {
			defer wg.Done()
			acc.Withdraw(1)
		}()
	}

	wg.Wait()

	// должен остаться 100
	// но при гонке может быть любое значение
	if acc.Balance != 100 {
		t.Errorf("Expected balance 100 after concurrent ops, got %d", acc.Balance)
	}
}
