/*
A hash set can store unique values, without any particular order.
*/
package hashset

import "github.com/zrcoder/dsgo"

// Set holds elements in go's native map
type Set[T comparable] struct {
	data map[T]dsgo.Empty
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{data: make(map[T]dsgo.Empty)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the dsgo.
func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.data[item] = dsgo.EmptyHolder
	}
}

// Remove removes the items (one or more) from the dsgo.
func (s *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s.data, item)
	}
}

// Contains check if items (one or more) are present in the dsgo.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty dsgo.
func (s *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, ok := s.data[item]; !ok {
			return false
		}
	}
	return true
}

// Len returns number of elements within the dsgo.
func (s *Set[T]) Len() int { return len(s.data) }

// Empty returns true if set does not contain any elements.
func (s *Set[T]) Empty() bool { return s.Len() == 0 }

// Values returns all items in the dsgo.
func (s *Set[T]) Values() []T {
	values := make([]T, s.Len())
	count := 0
	for item := range s.data {
		values[count] = item
		count++
	}
	return values
}

// Clear clears all values in the dsgo.
func (s *Set[T]) Clear() { clear(s.data) }

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "another".
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (s *Set[T]) Intersection(another *Set[T]) *Set[T] {
	res := New[T]()
	// Iterate over smaller set (optimization)
	if s.Len() <= another.Len() {
		s, another = another, s
	}
	for item := range s.data {
		if _, ok := another.data[item]; ok {
			res.Add(item)
		}
	}
	return res
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (s *Set[T]) Union(another *Set[T]) *Set[T] {
	res := New[T]()
	for item := range s.data {
		res.Add(item)
	}
	for item := range another.data {
		res.Add(item)
	}
	return res
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "another".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (s *Set[T]) Difference(another *Set[T]) *Set[T] {
	res := New[T]()
	for item := range s.data {
		if _, ok := another.data[item]; !ok {
			res.Add(item)
		}
	}
	return res
}
