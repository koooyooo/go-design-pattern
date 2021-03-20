package synchronized

import (
	"strings"
	"sync"
	"time"
)

type account struct {
	m       sync.Mutex
	w       *sync.WaitGroup // タイミング制御に使用、本質的な要素ではない
	balance int
	logs    []string
}

func NewAccount(w *sync.WaitGroup, balance int) *account {
	return &account{
		w:       w,
		balance: balance,
	}
}

func (a *account) PlusMinusTransaction(plusAmount, minusAmount int) {
	a.m.Lock()
	a.balance += plusAmount
	a.logs = append(a.logs, "Plus")
	time.Sleep(10 * time.Millisecond)
	a.balance -= minusAmount
	a.logs = append(a.logs, "Minus")
	a.w.Done()
	a.m.Unlock()
}

func (a account) String() string {
	return strings.Join(a.logs, ",")
}

func (a account) Amount() int {
	return a.balance
}
