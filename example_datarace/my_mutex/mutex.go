package my_mutex

type MyMutex struct {
	ch chan struct{}
}

// NewMyMutex создает новый мьютекс
func NewMyMutex() *MyMutex {
	m := &MyMutex{
		ch: make(chan struct{}, 1), // буфер 1 — бинарный семафор
	}
	m.ch <- struct{}{} // изначально "открыт"
	return m
}

// Lock блокирует мьютекс. Если уже заблокирован, ждёт.
func (m *MyMutex) Lock() {
	<-m.ch // получаем из канала — блокирует если пусто
}

// Unlock разблокирует мьютекс.
func (m *MyMutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked MyMutex")
	}
}
