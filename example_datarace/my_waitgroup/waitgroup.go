package my_waitgroup

import (
	"awesomeDataRace/example_datarace/my_mutex"
)

type MyWaitGroup struct {
	mu      *my_mutex.MyMutex
	counter int
	done    chan struct{}
}

// NewMyWaitGroup создает новый MyWaitGroup
func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{
		done: make(chan struct{}),
	}
}

// Add увеличивает счётчик горутин
func (wg *MyWaitGroup) Add(n int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	if wg.counter == 0 && n > 0 {
		// перезапускаем done, если нужен новый цикл ожидания
		wg.done = make(chan struct{})
	}

	wg.counter += n
	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}
}

// Done уменьшает счётчик на 1
func (wg *MyWaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter--
	if wg.counter < 0 {
		panic("negative WaitGroup counter")
	}
	if wg.counter == 0 {
		close(wg.done) // сигнал завершения
	}
}

// Wait блокирует до тех пор, пока счётчик не станет 0
func (wg *MyWaitGroup) Wait() {
	<-wg.done
}
