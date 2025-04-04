package main

import "sync"

// Примитивы
func RacePrimitive() int {
	var wg sync.WaitGroup
	var count int
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
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

func RaceObject() int {
	var wg sync.WaitGroup
	counter := &Counter{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
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

func RaceComposition() int {
	var wg sync.WaitGroup
	a := &A{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
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

func RaceAggregation() int {
	var wg sync.WaitGroup
	b := &B{}
	c := &C{b: b}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.b.value++
			}
		}()
	}
	wg.Wait()
	return c.b.value
}

// Ссылочные типы (срез)
func RaceSlice() []int {
	var wg sync.WaitGroup
	slice := make([]int, 0)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				slice = append(slice, j)
			}
		}()
	}
	wg.Wait()
	return slice
}
