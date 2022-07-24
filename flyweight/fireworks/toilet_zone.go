package fireworks

import "sync"

// eventToiletZone はイベントのトイレ区域を表す構造体
type eventToiletZone struct {
	m       sync.Mutex
	pool    sync.Pool
	Company *eventCompany
}

// NewEventToiletZone はイベントのトイレ区域を作成する
func NewEventToiletZone() *eventToiletZone {
	company := &eventCompany{}
	var m sync.Mutex
	var pool = sync.Pool{
		New: func() interface{} {
			m.Lock()
			defer m.Unlock()
			return company.BorrowToilet()
		},
	}
	return &eventToiletZone{
		m:       m,
		pool:    pool,
		Company: company,
	}
}

func (e *eventToiletZone) ReadyToilet(num int) {
	for i := 0; i < num; i++ {
		e.pool.Put(e.pool.New())
	}
}

func (e *eventToiletZone) AssignToilet() {
	v, _ := e.pool.Get().(*tmpToilet)
	v.Use()
	e.pool.Put(v)
}
