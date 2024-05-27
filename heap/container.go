package heap

import (
	"maps"
	"slices"

	"github.com/zrcoder/dsgo"
)

var _ dsgo.Container[int] = (*Heap[int])(nil)

// Len returns the size of the heap.
// The complexity is O(1)
func (h *Heap[T]) Len() int {
	if h.advanced {
		return h.size
	}
	return len(h.data)
}

// Empty returns if the heap is empty.
// The complexity is O(1)
func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}

// Values returns the sorted values in the heap
func (h *Heap[T]) Values() []T {
	res := make([]T, h.Len())
	tmp := h.clone()
	for i := range res {
		res[i], _ = tmp.Pop()
	}
	return res
}

func (h *Heap[T]) Clear() {
	clear(h.data)
	h.data = h.data[:0]
	clear(h.idx)
	clear(h.cnt)
	h.size = 0
}

func (h *Heap[T]) clone() *Heap[T] {
	return &Heap[T]{
		cmp:      h.cmp,
		data:     slices.Clone(h.data),
		idx:      maps.Clone(h.idx),
		cnt:      maps.Clone(h.cnt),
		size:     h.size,
		advanced: h.advanced,
	}
}
