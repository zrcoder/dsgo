/*
A stack gives you a LIFO or last-in first-out order.
You can only Push to add a new element to the top of the stack,
Pop to remove the element from the top,
and Peek at the top element without poping it off.
*/
package liststack

import (
	"github.com/zrcoder/dsgo/list"
	"github.com/zrcoder/dsgo/stack"
)

// Stack is a LIFO (last-in first-out) list
type Stack[T any] struct {
	list *list.List[T]
}

var _ stack.Stack[int] = (*Stack[int])(nil)

func New[T any]() *Stack[T] {
	return &Stack[T]{list.New[T]()}
}

// Push add a new element to the top
func (s *Stack[T]) Push(item T) {
	s.list.PushBack(item)
}

// Pop remove the element from the top and returns it
func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.list.Empty() {
		return
	}
	return s.list.Remove(s.list.Back()), true
}

// Peek returns the element from the top without deletion
func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.list.Empty() {
		return
	}
	return s.list.Back().Value, true
}
