package whitespace

type Stack[T any] struct {
	data []T
	bottom int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}, bottom: 0}
}

func (s *Stack[T]) Push(v T) {
	if s.bottom > len(s.data) {
		panic("stack bottom larger than stack")
	}

	if len(s.data) == s.bottom {
		s.data = append(s.data, v)
		s.bottom++
		return
	}

	s.data[s.bottom] = v
	s.bottom++
}

func (s *Stack[T]) Pop() T {
	if s.bottom <= 0 {
		panic("empty on stack.Pop()")
	}

	s.bottom--
	return s.data[s.bottom]
}

func (s *Stack[T]) Get(i int64) T {
	if i < 0 {
		panic("negative value in stack.Get()")
	}

	idx := int64(s.bottom) - i - 1
	if idx < 0 {
		panic("too deep in stack.Get()")
	}

	return s.data[idx]
}
