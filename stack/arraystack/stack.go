package arraystack

import (
	"slices"
)

// Stack holds elements in an array-list
type Stack[T comparable] struct {
	data []T
}

// New instantiates a new empty stack
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a value onto the top of the stack
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (s *Stack[T]) Pop() (value T, ok bool) {
	if len(s.data) == 0 {
		return
	}
	n := len(s.data)
	x := s.data[n-1]
	s.data = s.data[:n-1]
	return x, true
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (s *Stack[T]) Peek() (value T, ok bool) {
	if len(s.data) == 0 {
		return
	}
	return s.data[len(s.data)-1], true
}

// Len returns number of elements within the stack.
func (s *Stack[T]) Len() int { return len(s.data) }

// Empty returns true if stack does not contain any elements.
func (s *Stack[T]) Empty() bool { return len(s.data) == 0 }

// Values returns all elements in the stack (LIFO order).
func (s *Stack[T]) Values() []T {
	res := slices.Clone(s.data)
	slices.Reverse(res)
	return res
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	clear(s.data)
	s.data = s.data[:0]
}
