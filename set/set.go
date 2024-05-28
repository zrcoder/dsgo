package set

import "github.com/zrcoder/dsgo"

type Empty struct{}

var EmptyHolder = Empty{}

type Set[T comparable] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	// Intersection(Set[T]) Set[T]
	// Union(Set[T]) Set[T]
	// Difference(Set[T]) Set[T]

	dsgo.Container[T]
	// Len() int
	// Empty() bool
	// Values() []T
	// Clear()
}
