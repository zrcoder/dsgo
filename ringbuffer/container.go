package ringbuffer

import "github.com/zrcoder/dsgo"

var _ dsgo.Container[int] = (*Buffer[int])(nil)

// Len returns number of elements within the buffer.
func (b *Buffer[T]) Len() int {
	return b.size
}

// Empty returns true if buffer does not contain any elements.
func (b *Buffer[T]) Empty() bool {
	return b.Len() == 0
}

// Values returns all elements in the buffer (FIFO order).
func (b *Buffer[T]) Values() []T {
	values := make([]T, b.Len())
	for i := range values {
		values[i] = b.values[(b.start+i)%b.maxSize]
	}
	return values
}

// Clear removes all elements from the buffer.
func (b *Buffer[T]) Clear() {
	b.values = make([]T, b.maxSize)
	b.start = 0
	b.end = 0
	b.full = false
	b.size = 0
}
