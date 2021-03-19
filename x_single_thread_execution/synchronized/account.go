package synchronized

import (
	"strings"
	"sync"
	"time"
)

type Account struct {
	m      sync.Mutex
	amount int
	logs   []string
}

func (s *Account) Transaction(w *sync.WaitGroup) {
	s.m.Lock()
	s.amount += 100
	s.logs = append(s.logs, "Plus")
	time.Sleep(10 * time.Millisecond)
	s.amount -= 100
	s.logs = append(s.logs, "Minus")
	w.Done()
	s.m.Unlock()
}

func (s Account) String() string {
	return strings.Join(s.logs, ",")
}

func (s Account) Amount() int {
	return s.amount
}
