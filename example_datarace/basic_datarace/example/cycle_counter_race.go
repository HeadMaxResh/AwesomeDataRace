package example

import (
	"awesomeDataRace/example_datarace/my_waitgroup"
)

// Гонка на счетчике цикла

// Примитивы
func RacePrimitive(goroutines, iterations int) int {
	var wg *my_waitgroup.MyWaitGroup
	var count int
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

// Объекты
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func RaceObject(goroutines, iterations int) int {
	var wg *my_waitgroup.MyWaitGroup
	counter := &Counter{}
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	return counter.value
}

// Композиция
type B struct {
	value int
}

type A struct {
	B // Композиция
}

func RaceComposition(goroutines, iterations int) int {
	var wg *my_waitgroup.MyWaitGroup
	a := &A{}
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				a.value++
			}
		}()
	}
	wg.Wait()
	return a.value
}

// Агрегация
type C struct {
	b *B
}

func RaceAggregation(goroutines, iterations int) int {
	var wg *my_waitgroup.MyWaitGroup
	b := &B{}
	c := &C{b: b}
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				c.b.value++
			}
		}()
	}
	wg.Wait()
	return c.b.value
}

// Ссылочные типы (срез)
func RaceSlice(goroutines, iterations int) []int {
	var wg *my_waitgroup.MyWaitGroup
	slice := make([]int, 0)
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				slice = append(slice, j)
			}
		}()
	}
	wg.Wait()
	return slice
}
