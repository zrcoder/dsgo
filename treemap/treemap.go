// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"cmp"
	"fmt"
	"strings"

	"github.com/zrcoder/dsgo"
	rbt "github.com/zrcoder/dsgo/redblacktree"
)

// Map holds the elements in a red-black tree
type Map[K comparable, V any] struct {
	tree *rbt.Tree[K, V]
}

// New instantiates a tree map with the built-in comparator for K
func New[K cmp.Ordered, V any]() *Map[K, V] {
	return NewWith[K, V](cmp.Compare[K])
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith[K comparable, V any](comparator dsgo.Comparator[K]) *Map[K, V] {
	return &Map[K, V]{tree: rbt.NewWith[K, V](comparator)}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	m.tree.Remove(key)
}

// Len returns number of elements in the map.
func (m *Map[K, V]) Len() int {
	return m.tree.Len()
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.tree.Empty()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns 0-value, 0-value, false if map is empty.
func (m *Map[K, V]) Min() (key K, value V, ok bool) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value, true
	}
	return
}

// Max returns the maximum key and its value from the tree map.
// Returns 0-value, 0-value, false if map is empty.
func (m *Map[K, V]) Max() (key K, value V, ok bool) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value, true
	}
	return
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the map is empty, or because
// all keys in the map are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V, ok bool) {
	if node, ok := m.tree.Floor(key); ok {
		return node.Key, node.Value, true
	}
	return
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the map is empty, or because
// all keys in the map are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V, ok bool) {
	if node, ok := m.tree.Ceiling(key); ok {
		return node.Key, node.Value, true
	}
	return
}

// Inorer travels the tree in-order with a handler.
func (m *Map[K, V]) Inorder(handler func(key K, value V)) {
	m.tree.Inorder(handler)
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	buf := strings.Builder{}
	buf.WriteString("TreeMap\nmap[")
	m.tree.Inorder(func(key K, value V) {
		buf.WriteString(fmt.Sprintf("%v:%v ", key, value))
	})
	buf.WriteString("]")
	return buf.String()
}
