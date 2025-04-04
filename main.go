package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	// Количество горутин и итераций
	goroutines := 2
	iterations := 1000

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				// Data race происходит здесь: несинхронизированное чтение и запись
				counter++
			}
		}()
	}

	wg.Wait()

	// Ожидаем 2000, но из-за гонки данных результат будет меньше
	fmt.Println("Итоговое значение:", counter)
}
