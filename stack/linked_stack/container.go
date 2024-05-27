package linked_stack

func (s *Stack[T]) Len() int {
	return s.size
}

func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) Values() []T {
	res := make([]T, s.size)
	p := s.top
	for i := range res {
		res[i] = p.value
		p = p.next
	}
	return res
}

func (s *Stack[T]) Clear() {
	s.top = nil
	s.size = 0
}
