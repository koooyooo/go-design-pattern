package un_synchronized

import (
	"strings"
	"sync"
	"time"
)

type UnSynchronized struct {
	amount int
	logs   []string
}

func (s *UnSynchronized) SinglePath(w *sync.WaitGroup) {
	s.amount += 100
	s.logs = append(s.logs, "Plus")
	time.Sleep(10 * time.Millisecond)
	s.amount -= 100
	s.logs = append(s.logs, "Minus")
	w.Done()
}

func (s UnSynchronized) String() string {
	return strings.Join(s.logs, ",")
}

func (s UnSynchronized) Amount() int {
	return s.amount
}
