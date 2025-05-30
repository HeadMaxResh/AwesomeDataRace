package example

import (
	"time"
)

type Inventory struct {
	Stock int
}

func (inv *Inventory) Purchase(quantity int) bool {
	if inv.Stock >= quantity {
		time.Sleep(1 * time.Millisecond)
		inv.Stock -= quantity
		return true
	}
	return false
}
