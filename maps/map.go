package maps

import "github.com/zrcoder/dsgo"

// Map interface that all maps implement
type Map[K comparable, V any] interface {
	Remove(key K)

	Cache[K, V]
}

type Cache[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Keys() []K

	dsgo.Container[V]
	// Len() int
	// Empty() bool
	// Values() []T
	// Clear()
}

// Pair holds a key-value item
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap[K comparable, V comparable] interface {
	GetKey(value V) (key K, found bool)

	Map[K, V]
}
