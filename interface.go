package dsgo

type Container[T any] interface {
	Len() int
	Empty() bool
	Values() []T
	Clear()
}

type Stack[T any] interface {
	Container[T]
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}

type Cache[K comparable, V any] interface {
	Container[V]
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Keys() []K
}

type Map[K comparable, V any] interface {
	Cache[K, V]
	Remove(key K)
}

type BidMap[K comparable, V comparable] interface {
	Map[K, V]
	GetKey(value V) (key K, found bool)
}

type Set[T comparable] interface {
	Container[T]
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool
	// Intersection(Set[T]) Set[T]
	// Union(Set[T]) Set[T]
	// Difference(Set[T]) Set[T]
}
