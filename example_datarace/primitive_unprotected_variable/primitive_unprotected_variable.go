package primitive_unprotected_variable

import "sync"

// Примитивная незащищенная переменная

func RacePrimitive(goroutines, iterations int) int {
	var wg sync.WaitGroup
	count := 0

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				count++
			}
		}()
	}

	wg.Wait()
	return count
}
