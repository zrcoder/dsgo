package linked_stack

import "github.com/zrcoder/dsgo/stack"

type node[T any] struct {
	value T
	next  *node[T]
}

// Stack is a LIFO (last-in first-out) list
type Stack[T any] struct {
	top  *node[T]
	size int
}

var _ stack.Stack[int] = (*Stack[int])(nil)

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

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
