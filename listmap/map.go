// Package listmap is a map that preserves insertion-order.
//
// It is backed by a buit-in hash map to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package listmap

import (
	"fmt"
	"strings"

	"github.com/zrcoder/dsgo"
	"github.com/zrcoder/dsgo/list"
)

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map[K comparable, V any] struct {
	m    map[K]*list.Element[dsgo.Pair[K, V]]
	list *list.List[dsgo.Pair[K, V]]
}

// New instantiates a linked-hash-map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:    make(map[K]*list.Element[dsgo.Pair[K, V]]),
		list: list.New[dsgo.Pair[K, V]](),
	}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	kv := dsgo.Pair[K, V]{Key: key, Value: value}
	if e, ok := m.m[key]; !ok {
		e := m.list.PushBack(kv)
		m.m[key] = e
	} else {
		e.Value = kv
	}
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	if e, ok := m.m[key]; ok {
		return e.Value.Value, true
	}
	return
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	if e, ok := m.m[key]; ok {
		delete(m.m, key)
		m.list.Remove(e)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.Len() == 0
}

// Len returns number of elements in the map.
func (m *Map[K, V]) Len() int {
	return m.list.Len()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.list.Len())
	m.Inorder(func(key K, value V) {
		keys = append(keys, key)
	})
	return keys
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.list.Len())
	m.Inorder(func(key K, value V) {
		values = append(values, value)
	})
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	clear(m.m)
	m.list.Clear()
}

// Inorder travels the map in inserted order
func (m *Map[K, V]) Inorder(hander func(key K, value V)) {
	for e := m.list.Front(); e != nil; e = e.Next() {
		hander(e.Value.Key, e.Value.Value)
	}
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "LinkedHashMap\nmap["
	kvs := make([]string, 0, m.list.Len())
	for e := m.list.Front(); e != nil; e = e.Next() {
		kvs = append(kvs, fmt.Sprintf("%v:%v", e.Value.Key, e.Value.Value))
	}
	return str + strings.Join(kvs, " ") + " ]"
}
