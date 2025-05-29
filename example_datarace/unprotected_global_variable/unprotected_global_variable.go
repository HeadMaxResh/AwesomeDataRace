package unprotected_global_variable

import "sync"

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
