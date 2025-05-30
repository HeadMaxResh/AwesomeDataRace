package example

import (
	"net"
	"sync"
)

// Незащищенная глобальная переменная

var globalCount int

func RaceGlobal(goroutines, iterations int) int {
	var wg sync.WaitGroup

	globalCount = 0

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				globalCount++
			}
		}()
	}

	wg.Wait()
	return globalCount
}

// пример из текста
var service map[string]net.Addr

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}
