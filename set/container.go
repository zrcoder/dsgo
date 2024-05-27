package set

import "github.com/zrcoder/dsgo"

// Assert Set implementation
var _ dsgo.Container[int] = (*Set[int])(nil)

// Len returns number of elements within the set.
func (set *Set[T]) Len() int {
	return len(set.data)
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) Empty() bool {
	return set.Len() == 0
}

// Values returns all items in the set.
func (set *Set[T]) Values() []T {
	values := make([]T, set.Len())
	count := 0
	for item := range set.data {
		values[count] = item
		count++
	}
	return values
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	clear(set.data)
}
