package liststack

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Empty() bool {
	return s.list.Empty()
}

func (s *Stack[T]) Values() []T {
	return s.list.Values()
}

func (s *Stack[T]) Clear() {
	s.list.Clear()
}
