/*
A set can store unique values, without any particular order.
*/
package set

// Set holds elements in go's native map
type Set[T comparable] struct {
	data map[T]empty
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{data: make(map[T]empty)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.data[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(set.data, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.data[item]; !contains {
			return false
		}
	}
	return true
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "another".
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set[T]) Intersection(another *Set[T]) *Set[T] {
	res := New[T]()
	// Iterate over smaller set (optimization)
	if set.Len() <= another.Len() {
		set, another = another, set
	}
	for item := range set.data {
		if _, contains := another.data[item]; contains {
			res.Add(item)
		}
	}
	return res
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "another" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[T]) Union(another *Set[T]) *Set[T] {
	res := New[T]()
	for item := range set.data {
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
func (set *Set[T]) Difference(another *Set[T]) *Set[T] {
	res := New[T]()
	for item := range set.data {
		if _, contains := another.data[item]; !contains {
			res.Add(item)
		}
	}
	return res
}
