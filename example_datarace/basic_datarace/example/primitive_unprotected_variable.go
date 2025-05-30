package example

import (
	"awesomeDataRace/example_datarace/my_waitgroup"
	"fmt"
	"os"
	"time"
)

// Примитивная незащищенная переменная

func RaceUnprotectedPrimitive(goroutines, iterations int) int {
	var wg *my_waitgroup.MyWaitGroup
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

// пример из текста
type Watchdog struct{ Last int64 }

func (w *Watchdog) KeepAlive() {
	// Первый конфликтующий доступ.
	w.Last = time.Now().UnixNano()
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Второй конфликтующий доступ.
			if w.Last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
