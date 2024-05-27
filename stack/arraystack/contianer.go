package arraystack

import (
	"slices"
)

// Len returns number of elements within the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// Empty returns true if stack does not contain any elements.
func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

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
