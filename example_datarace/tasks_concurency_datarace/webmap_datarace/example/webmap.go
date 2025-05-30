package example

import (
	"fmt"
	"sync"
)

var pageViewsMap = make(map[string]int)

var lock sync.RWMutex

func IncrementPageView(page string) {
	lock.Lock()
	pageViewsMap[page]++
	lock.Unlock()
}

func GetAllPageViews() map[string]int {
	lock.RLock()
	defer lock.RUnlock()

	copiedMap := make(map[string]int)
	for k, v := range pageViewsMap {
		copiedMap[k] = v
	}
	return copiedMap
}

func ResetPageViewsMap() {
	lock.Lock()
	defer lock.Unlock()

	pageViewsMap = make(map[string]int)
}

func ShowPageViews() {
	lock.RLock()
	for page, views := range pageViewsMap {
		fmt.Printf("%s: %d просмотров\n", page, views)
	}
	lock.RUnlock()
}

func simulateActivity(page string, iterations int) {
	for i := 0; i < iterations; i++ {
		IncrementPageView(page)
	}
}

func main() {
	pages := []string{"Главная страница", "Контакты", "Продукты"}

	var waitGroup sync.WaitGroup

	for _, page := range pages {
		waitGroup.Add(1)
		go func(p string) {
			simulateActivity(p, 1000)
			waitGroup.Done()
		}(page)
	}

	waitGroup.Wait()

	ShowPageViews()
}
