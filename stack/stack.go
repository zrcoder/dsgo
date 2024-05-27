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
