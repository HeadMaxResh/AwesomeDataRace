package test

import (
	"awesomeDataRace/example_datarace/tasks_concurency_datarace/database_datarace/example"
	"math/rand"
	"sync"
	"testing"
)

func TestUsersDBRaceCondition(t *testing.T) {
	db := example.NewUsersDB()

	const usersCount = 10
	for i := 1; i <= usersCount; i++ {
		err := db.AddUser(i)
		if err != nil {
			t.Fatalf("Ошибка добавления пользователя: %v", err)
		}
	}

	var wg sync.WaitGroup
	numWorkers := 100
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(int64(workerId)))
			for j := 0; j < 100; j++ {
				id := rnd.Intn(usersCount) + 1
				delta := int64(rnd.Intn(10))
				err := db.UpdateScore(id, delta)
				if err != nil {
					t.Errorf("Рабочий поток #%d: ошибка обновления счета для пользователя %d: %s", workerId, id, err)
				}
			}
		}(i)
	}

	wg.Wait()

	totalSum := db.SumScores()
	expectedTotalSum := int64(0)
	for i := 1; i <= usersCount; i++ {
		user, _ := db.GetUser(i)
		expectedTotalSum += user.Score
	}

	if totalSum != expectedTotalSum {
		t.Errorf("Сумма баллов неверна! Ожидаемая общая сумма: %d, реальная сумма: %d", expectedTotalSum, totalSum)
	}

	db.PrintScores()
}
