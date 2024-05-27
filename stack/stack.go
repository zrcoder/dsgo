/*
A stack gives you a LIFO or last-in first-out order.
You can only Push to add a new element to the top of the stack,
Pop to remove the element from the top,
and Peek at the top element without poping it off.
*/
package stack

import "github.com/zrcoder/dsgo"

// Stack interface that all stacks implement
type Stack[T any] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	dsgo.Container[T]
	// Len() int
	// Empty() bool
	// Values() []T
	// Clear()
}
