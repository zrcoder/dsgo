package heap

import (
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

func (h *Heap[T]) Empty() bool {
	return h.size == 0
}

func (h *Heap[T]) Values() []T {
	if !h.advanced {
		return slices.Clone(h.data)
	}
	res := make([]T, 0, h.size)
	for _, v := range h.data {
		for n := h.cnt[v]; n > 0; n-- {
			res = append(res, v)
		}
	}
	return res
}

func (h *Heap[T]) Clear() {
	clear(h.data)
	h.data = h.data[:0]
}
