package test

import (
	"awesomeDataRace/example_datarace/inventory_datarace/example"
	"sync"
	"testing"
)

func TestSafeInventory_ConcurrentFailWhenNotEnoughStock(t *testing.T) {
	inv := &example.SafeInventory{Stock: 5}
	var wg sync.WaitGroup
	var successCount int
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if inv.Purchase(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	finalStock := inv.GetStock()
	if finalStock < 0 {
		t.Errorf("Stock went negative! Final stock: %d", finalStock)
	}
	if successCount > 5 {
		t.Errorf("Too many successful purchases: %d (stock was only 5)", successCount)
	}
}

func TestSafeInventory_ConcurrentPurchases(t *testing.T) {
	inv := &example.SafeInventory{Stock: 100}
	var wg sync.WaitGroup
	var mu sync.Mutex
	success := 0

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if inv.Purchase(3) {
				mu.Lock()
				success++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	expected := 100 - inv.GetStock()

	if expected != success*3 {
		t.Errorf("Mismatch: %d units reported purchased, actual stock drop: %d", success*3, expected)
	}
	if inv.GetStock() < 0 {
		t.Errorf("Stock went negative: %d", inv.GetStock())
	} else {
		t.Logf("Final stock (safe): %d, successful purchases: %d", inv.GetStock(), success)
	}
}
