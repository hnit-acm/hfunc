package structure

import "sync"

type Stack struct {
	mu   sync.RWMutex
	data *Array
	size int
}

// 得到栈顶
func (s *Stack) Peek() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.IsEmpty() {
		return nil
	}
	return s.data.Get(s.size - 1)
}

// 弹出栈顶
func (s *Stack) Poll() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsEmpty() {
		return nil
	}
	s.size--
	return s.data.Get(s.size)
}

// 入栈
func (s *Stack) Push(target interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.data == nil {
		s.data = NewDefaultArray()
	}
	s.data.Set(s.size, target)
	s.size++
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Len() int {
	return s.size
}
