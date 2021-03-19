package synchronized

import (
	"strings"
	"sync"
	"time"
)

type Synchronized struct {
	m      sync.Mutex
	amount int
	logs   []string
}

func (s *Synchronized) SinglePath(w *sync.WaitGroup) {
	s.m.Lock()
	s.amount += 100
	s.logs = append(s.logs, "Plus")
	time.Sleep(10 * time.Millisecond)
	s.amount -= 100
	s.logs = append(s.logs, "Minus")
	w.Done()
	s.m.Unlock()
}

func (s Synchronized) String() string {
	return strings.Join(s.logs, ",")
}

func (s Synchronized) Amount() int {
	return s.amount
}
