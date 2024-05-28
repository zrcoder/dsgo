package maps

import "github.com/zrcoder/dsgo"

// Map interface that all maps implement
type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	dsgo.Container[V]
	// Len() int
	// Empty() bool
	// Values() []T
	// Clear()
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap[K comparable, V comparable] interface {
	GetKey(value V) (key K, found bool)

	Map[K, V]
}
