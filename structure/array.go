package structure

type Array struct {
	elements []interface{}
	length   int
}

const (
	defaultCapacity = 16
	expand          = 2
)

func (s *Array) Len() int {
	return s.length
}

func (s *Array) Cap() int {
	return len(s.elements)
}
func (s *Array) Get(index int) interface{} {
	if index >= s.length {
		return nil
	}
	return s.elements[index]
}

func (s *Array) Set(index int, element interface{}) interface{} {
	if index > s.length {
		return nil
	}
	if index == s.length {
		s.add(element)
	}
	s.elements[index] = element
	return element
}

func (s *Array) add(element interface{}) {
	s.ensureCapacity()
	s.elements[s.length] = element
	s.length++
}

func (s *Array) ensureCapacity() {
	if s.length < len(s.elements) {
		return
	}
	temp := s.elements
	s.elements = make([]interface{}, s.length*expand, s.length*expand)
	for i, x := range temp {
		s.elements[i] = x
	}
}

func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		elements: make([]interface{}, capacity, capacity),
		length:   0,
	}
}

func NewDefaultArray() *Array {
	return NewArray(defaultCapacity)
}
