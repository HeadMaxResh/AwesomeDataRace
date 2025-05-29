package randomly_shared_variable

import "sync"

// Случайно разделяемая переменная

func RaceAccidentalSharing(goroutines, iterations int) int {
	var wg sync.WaitGroup
	var count int

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				count++
			}
		}(i)
	}

	wg.Wait()
	return count
}
