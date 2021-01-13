package structure

import "sync"

// todo 自动扩容队列的实现
// assignees: yinrenxin,nieaowei,leoiwan
// labels: enhancement

type Queue struct {
	mu           sync.RWMutex
	data         *Array
	header, tail int
}

// 得到队首
func (q *Queue) Peek() interface{} {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if q.IsEmpty() {
		return nil
	}
	return q.data.Get(q.header)
}

// 弹出队首
func (q *Queue) Poll() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.IsEmpty() {
		return nil
	}
	result := q.data.Get(q.header)
	q.header++
	if q.header*2 >= q.data.Cap() {
		q.move()
	}
	return result
}

// 入队
func (q *Queue) Push(target interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.data == nil {
		q.data = NewDefaultArray()
	}
	q.data.Set(q.tail, target)
	q.tail++
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Len() int {
	return q.tail - q.header
}

/**
如果队首距离数组首部过远，将队首拉到数组首部
防止不断扩容
*/
func (q *Queue) move() {
	for i := 0; i < q.Len(); i++ {
		q.data.Set(i, q.data.Get(i+q.header))
	}
	q.tail = q.Len()
	q.header = 0
}
