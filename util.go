package dsgo

type Empty struct{}

var EmptyHolder = Empty{}

// Pair holds a key-value item
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}
