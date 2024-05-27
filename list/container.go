package list

import "github.com/zrcoder/dsgo"

var _ dsgo.Container[int] = (*List[int])(nil)

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[T]) Len() int { return l.len }

// Empty returns if the list is empty.
// The complexity is O(1).
func (l *List[T]) Empty() bool { return l.len == 0 }

// Values returns the values slice in the list
func (l *List[T]) Values() []T {
	res := make([]T, l.len)
	for i, e := 0, l.Front(); i < len(res); i, e = i+1, e.Next() {
		res[i] = e.Value
	}
	return res
}

func (l *List[T]) Clear() {
	l.Init()
}
