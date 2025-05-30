package example

import (
	"awesomeDataRace/example_datarace/my_waitgroup"
	"os"
)

// Случайно разделяемая переменная

func RaceAccidentalSharing(goroutines, iterations int) int {
	var wg *my_waitgroup.MyWaitGroup
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

// пример из текста
// ParallelWrite записывает данные в file1 и file2,
// возвращает ошибки.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// Эта err разделяемая с main goroutine,
			// поэтому выполнение записи вызывает
			// гонку с выполнением записи ниже.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}

	// Вторая конфликтующая запись в err.
	f2, err := os.Create("file2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
