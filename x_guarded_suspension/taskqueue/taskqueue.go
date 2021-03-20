package taskqueue

import "sync"

type taskQueue struct {
	cond  *sync.Cond
	tasks []interface{}
	logs  []string
}

func NewTaskQueue() *taskQueue {
	return &taskQueue{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (q *taskQueue) AddLast(task interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.tasks = append(q.tasks, task)
	q.logs = append(q.logs, "AddLast")
	q.cond.Signal()
}

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

func (q taskQueue) Logs() []string {
	return q.logs
}
