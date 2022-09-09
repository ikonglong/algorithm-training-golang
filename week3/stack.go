package week3

type Stack struct {
	items []interface{}
	n     int
}

func NewStack(initCap int) *Stack {
	return &Stack{
		items: make([]interface{}, 0, initCap),
		n:     0,
	}
}

func (s *Stack) push(v interface{}) {
	s.items = append(s.items, v)
	s.n++
}

func (s *Stack) pop() interface{} {
	top := s.items[s.n-1]
	s.items = s.items[:s.n-1]
	s.n--
	return top
}

func (s *Stack) peek() interface{} {
	if s.isEmpty() {
		return nil
	}
	return s.items[s.n-1]
}

func (s *Stack) len() int {
	return s.n
}

func (s *Stack) isEmpty() bool {
	return s.n == 0
}
