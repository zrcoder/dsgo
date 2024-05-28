package linkedstack

type node[T any] struct {
	value T
	next  *node[T]
}

// Stack is a LIFO (last-in first-out) list
type Stack[T any] struct {
	top  *node[T]
	size int
}

func New[T any]() *Stack[T] { return &Stack[T]{} }

// Push add a new element to the top
func (s *Stack[T]) Push(value T) {
	node := &node[T]{value: value, next: s.top}
	s.top = node
	s.size++
}

// Pop remove the element from the top and returns it
func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.Empty() {
		return
	}
	res := s.top
	s.top = res.next
	res.next = nil
	s.size--
	return res.value, true
}

// Peek returns the element from the top without deletion
func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.Empty() {
		return
	}
	return s.top.value, true
}

func (s *Stack[T]) Len() int { return s.size }

func (s *Stack[T]) Empty() bool { return s.size == 0 }

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
