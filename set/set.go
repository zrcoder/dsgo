/*
A set can store unique values, without any particular order.
*/
package set

import "github.com/zrcoder/dsgo"

type Set[T comparable] map[T]dsgo.Empty

func New[T comparable]() Set[T] {
	return make(map[T]dsgo.Empty)
}

func NewWithCapacity[T comparable](c int) Set[T] {
	return make(map[T]dsgo.Empty, c)
}

func (s Set[T]) Add(item T) {
	s[item] = dsgo.Empty{}
}

func (s Set[T]) Delete(item T) {
	delete(s, item)
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Values() []T {
	r := make([]T, s.Size())
	i := 0
	for key := range s {
		r[i] = key
		i++
	}
	return r
}

func (s Set[T]) Range(f func(item T) bool) {
	for key := range s {
		if f(key) {
			return
		}
	}
}
