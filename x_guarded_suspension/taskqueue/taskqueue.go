package taskqueue

import "sync"

// taskQueue は同期制御機能を持ったQueueです
type taskQueue struct {
	cond  *sync.Cond
	tasks []interface{}
	logs  []string
}

// NewTaskQueue は新規のtaskQueueを生成します
func NewTaskQueue() *taskQueue {
	return &taskQueue{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

// AddLast はtaskQueueの最後に要素を追加します
func (q *taskQueue) AddLast(task interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.tasks = append(q.tasks, task)
	q.logs = append(q.logs, "AddLast")
	q.cond.Signal()
}

// RemoveFirst はtaskQueueの先頭の要素を削除しそれを返します
func (q *taskQueue) RemoveFirst() interface{} {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.tasks) == 0 {
		q.logs = append(q.logs, "RemoveFirst-Wait")
		q.cond.Wait()
	}
	v := q.tasks[0]
	q.tasks = q.tasks[1:]
	q.logs = append(q.logs, "RemoveFirst-Done")
	return v
}

// Logs はログを返します
func (q taskQueue) Logs() []string {
	return q.logs
}
