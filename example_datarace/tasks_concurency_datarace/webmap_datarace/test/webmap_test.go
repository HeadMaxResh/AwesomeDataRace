package test

import (
	"awesomeDataRace/example_datarace/tasks_concurency_datarace/webmap_datarace/example"
	"sync"
	"testing"
)

func TestConcurrentAccess(t *testing.T) {
	example.ResetPageViewsMap()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			example.IncrementPageView("страница А")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			example.IncrementPageView("страница B")
		}
		wg.Done()
	}()

	wg.Wait()

	allPages := example.GetAllPageViews()

	actualA := allPages["страница А"]
	actualB := allPages["страница B"]

	if actualA != 1000 || actualB != 1000 {
		t.Errorf("Некорректные значения: A=%d, B=%d", actualA, actualB)
	}

	example.ShowPageViews()
}
