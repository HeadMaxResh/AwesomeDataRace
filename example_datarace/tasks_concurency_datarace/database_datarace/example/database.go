package example

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	ID    int
	Score int64
}

type UsersDB struct {
	users map[int]*User
	lock  sync.RWMutex
}

func NewUsersDB() *UsersDB {
	return &UsersDB{
		users: make(map[int]*User),
	}
}

func (db *UsersDB) AddUser(id int) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, exists := db.users[id]; exists {
		return fmt.Errorf("Пользователь с id %d уже существует", id)
	}
	user := &User{ID: id, Score: 0}
	db.users[id] = user
	return nil
}

func (db *UsersDB) GetUser(id int) (*User, bool) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	user, found := db.users[id]
	return user, found
}

func (db *UsersDB) UpdateScore(id int, delta int64) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	user, found := db.users[id]
	if !found {
		return fmt.Errorf("Пользователь с id %d не найден", id)
	}
	user.Score += delta
	return nil
}

func (db *UsersDB) PrintScores() {
	db.lock.RLock()
	defer db.lock.RUnlock()

	for _, user := range db.users {
		fmt.Printf("ID: %d | Баллов: %d\n", user.ID, user.Score)
	}
}

func (db *UsersDB) SimulateRaces(numWorkers int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numWorkers; i++ {
		go func(workerId int) {
			for j := 0; j < 100; j++ {
				id := rand.Intn(len(db.users)) + 1
				delta := int64(rand.Intn(10))
				err := db.UpdateScore(id, delta)
				if err != nil {
					fmt.Printf("Worker #%d: Error updating score for user %d: %s\n", workerId, id, err)
				}
			}
		}(i)
	}
}

func (db *UsersDB) SumScores() int64 {
	db.lock.RLock()
	defer db.lock.RUnlock()

	sum := int64(0)
	for _, user := range db.users {
		sum += user.Score
	}
	return sum
}
