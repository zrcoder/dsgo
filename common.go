package dsgo

import "cmp"

type Comparator[T any] func(a, b T) int

func OrderedComparator[T cmp.Ordered]() Comparator[T] {
	return func(a, b T) int {
		return cmp.Compare(a, b)
	}
}

func (cmp Comparator[T]) Reverse() Comparator[T] {
	return func(a, b T) int {
		return -cmp(a, b)
	}
}

type Container[T any] interface {
	Len() int    // size of the container
	Empty() bool // if the container is empty
	Values() []T // values the container holds
	Clear()      // clears the container
}
