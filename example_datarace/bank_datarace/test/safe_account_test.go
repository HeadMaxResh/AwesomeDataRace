package test

import (
	"awesomeDataRace/example_datarace/bank_datarace/example"
	"sync"
	"testing"
)

func TestSafeAccount_WithdrawAndDepositRace(t *testing.T) {
	acc := &example.SafeAccount{ID: 1, Balance: 1000}
	var wg sync.WaitGroup

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

	balance := acc.GetBalance()
	if balance != 1000 {
		t.Errorf("Expected balance 1000, got %d", balance)
	}
}

func TestSafeAccount_MultipleWithdrawals(t *testing.T) {
	acc := &example.SafeAccount{ID: 1, Balance: 100}
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

	totalWithdrawn := 100 - acc.GetBalance()
	if totalWithdrawn > 100 {
		t.Errorf("Overdrawn: total withdrawn = %d, balance = %d", totalWithdrawn, acc.GetBalance())
	} else {
		t.Logf("Withdrawn: %d, Failed ops: %d, Final balance: %d", totalWithdrawn, failures, acc.GetBalance())
	}
}

func TestSafeAccount_ConcurrentDeposits(t *testing.T) {
	acc := &example.SafeAccount{ID: 1, Balance: 0}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			acc.Deposit(1)
		}()
	}
	wg.Wait()

	if acc.GetBalance() != 1000 {
		t.Errorf("Incorrect final balance: expected 1000, got %d", acc.GetBalance())
	}
}

func TestSafeAccount_TransferRace(t *testing.T) {
	acc1 := &example.SafeAccount{ID: 1, Balance: 1000}
	acc2 := &example.SafeAccount{ID: 2, Balance: 1000}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			example.SafeTransfer(acc1, acc2, 10)
		}()
		go func() {
			defer wg.Done()
			example.SafeTransfer(acc2, acc1, 10)
		}()
	}

	wg.Wait()

	total := acc1.GetBalance() + acc2.GetBalance()
	if total != 2000 {
		t.Errorf("Inconsistent total: expected 2000, got %d", total)
	} else {
		t.Logf("Acc1: %d, Acc2: %d, Total: %d", acc1.GetBalance(), acc2.GetBalance(), total)
	}
}

func TestSafeAccount_DataConsistency(t *testing.T) {
	acc := &example.SafeAccount{ID: 1, Balance: 100}
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

	balance := acc.GetBalance()
	if balance != 100 {
		t.Errorf("Expected balance 100 after concurrent ops, got %d", balance)
	}
}
