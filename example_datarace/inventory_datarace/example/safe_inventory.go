package example

import (
	"sync"
	"time"
)

type SafeInventory struct {
	Stock int
	mu    sync.Mutex
}

func (inv *SafeInventory) Purchase(quantity int) bool {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	if inv.Stock >= quantity {
		time.Sleep(1 * time.Millisecond)
		inv.Stock -= quantity
		return true
	}
	return false
}

func (inv *SafeInventory) GetStock() int {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	return inv.Stock
}
