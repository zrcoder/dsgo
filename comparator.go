package dsgo

type Comparator[T any] func(a, b T) int

func Reverse[T any](cmp Comparator[T]) Comparator[T] {
	return func(a, b T) int {
		return -cmp(a, b)
	}
}
