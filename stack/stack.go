/*
A stack gives you a LIFO or last-in first-out order.
You can only Push to add a new element to the top of the stack,
Pop to remove the element from the top,
and Peek at the top element without poping it off.
*/
package stack

import "github.com/zrcoder/dsgo/list"

// Stack is a LIFO (last-in first-out) list
type Stack[T any] struct {
	*list.List[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{list.New[T]()}
}

// Push add a new element to the top
func (s *Stack[T]) Push(item T) {
	s.List.PushBack(item)
}

// Pop remove the element from the top and returns it
func (s *Stack[T]) Pop() T {
	return s.Remove(s.Back())
}

// Peek returns the element from the top without deletion
func (s *Stack[T]) Peek() T {
	return s.Back().Value
}
