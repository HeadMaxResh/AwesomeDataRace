package test

import (
	"awesomeDataRace/example_datarace/inventory_datarace/example"
	"sync"
	"testing"
)

func TestInventory_ConcurrentFailWhenNotEnoughStock(t *testing.T) {
	inv := &example.Inventory{Stock: 5}
	var wg sync.WaitGroup
	var successCount int
	var mu sync.Mutex

	// 10 клиентов пытаются купить по 1 (доступно только 5)
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

	if inv.Stock < 0 {
		t.Errorf("Stock went negative! Final stock: %d", inv.Stock)
	}
	if successCount > 5 {
		t.Errorf("Too many successful purchases: %d (stock was only 5)", successCount)
	}
}

func TestInventory_ConcurrentPurchasesRace(t *testing.T) {
	inv := &example.Inventory{Stock: 100}
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inv.Purchase(3)
		}()
	}

	wg.Wait()

	if inv.Stock < 0 {
		t.Errorf("Race condition: stock went negative: %d", inv.Stock)
	} else {
		t.Logf("Final stock (unsafe): %d", inv.Stock)
	}
}
