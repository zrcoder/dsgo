package heapx

import "github.com/zrcoder/dsgo/internal/heap"

// Remove removes value (any element) from the heap.
// This method is only supported for advanced Heap, it will panics if called on a base heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Remove(value T) bool {
	if h.cnt[value] == 0 {
		return false
	}
	return h.removeIndex(h.idx[value])
}

// Update re-establishes the heap ordering after the element has changed its value.
// This method is only supported for advanced Heap, it will panics if called on a base heap.
// Changing the value of the element value and then calling Update is equivalent to,
// but less expensive than, calling Remove(value) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Update(value T) bool {
	if h.cnt[value] == 0 {
		return false
	}
	heap.FixIndex(h, h.idx[value])
	return true
}

func (h *Heap[T]) removeIndex(index int) bool {
	if index < 0 || index >= len(h.data) {
		return false
	}
	value := h.data[index]
	if h.cnt[value] == 1 {
		heap.RemoveIndex(h, index)
		delete(h.idx, value)
	}
	h.size--
	h.cnt[value]--
	return true
}
