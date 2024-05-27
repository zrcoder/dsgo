package dsgo

import "cmp"

type (
	Empty struct{}

	Comparator[T any] func(a, b T) bool
)

func OrderedComparator[T cmp.Ordered]() Comparator[T] {
	return func(a, b T) bool {
		return cmp.Compare(a, b) < 0
	}
}

func (cmp Comparator[T]) Reverse() Comparator[T] {
	return func(a, b T) bool {
		return !cmp(a, b)
	}
}
