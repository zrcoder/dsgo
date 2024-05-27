package ringbuffer

// Buffer holds values in a slice.
type Buffer[T comparable] struct {
	values  []T
	start   int
	end     int
	full    bool
	maxSize int
	size    int
}

// New instantiates a new empty buffer with the specified size of maximum number of elements that it can hold.
// This max size of the buffer cannot be changed.
func New[T comparable](maxSize int) *Buffer[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}
	buffer := &Buffer[T]{maxSize: maxSize}
	buffer.Clear()
	return buffer
}

// Enqueue adds a value to the end of the buffer
func (b *Buffer[T]) Enqueue(value T) {
	if b.Full() {
		b.Dequeue()
	}
	b.values[b.end] = value
	b.end = (b.end + 1) % b.maxSize
	if b.end == b.start {
		b.full = true
	}

	b.size = b.calculateSize()
}

// Dequeue removes first element of the buffer and returns it, or the 0-value if buffer is empty.
// Second return parameter is true, unless the buffer was empty and there was nothing to pop.
func (b *Buffer[T]) Dequeue() (value T, ok bool) {
	if b.Empty() {
		return value, false
	}
	value, ok = b.values[b.start], true
	b.start = (b.start + 1) % b.maxSize
	b.full = false
	b.size = b.size - 1

	return
}

// Peek returns first element of the buffer without removing it, or nil if buffer is empty.
// Second return parameter is true, unless the buffer was empty and there was nothing to peek.
func (b *Buffer[T]) Peek() (value T, ok bool) {
	if b.Empty() {
		return value, false
	}
	return b.values[b.start], true
}

// Full returns true if the buffer is full, i.e. has reached the maximum number of elements that it can hold.
func (b *Buffer[T]) Full() bool {
	return b.Len() == b.maxSize
}

// Get gets the value at the given index index.
func (b *Buffer[T]) Get(index int) (value T, ok bool) {
	if b.withinRange(index) {
		return b.values[(b.start+index)%b.maxSize], true
	}
	return
}

// Set sets the value at the given index.
func (b *Buffer[T]) Set(index int, value T) bool {
	if b.withinRange(index) {
		b.values[(b.start+index)%b.maxSize] = value
		return true
	}
	return false
}

// Check that the index is within bounds of the list
func (b *Buffer[T]) withinRange(index int) bool {
	return index >= 0 && index < b.size
}

func (b *Buffer[T]) calculateSize() int {
	if b.end < b.start {
		return b.maxSize - b.start + b.end
	} else if b.end == b.start {
		if b.full {
			return b.maxSize
		}
		return 0
	}
	return b.end - b.start
}
