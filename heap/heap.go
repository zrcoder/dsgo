package heap

import (
	"cmp"

	"github.com/zrcoder/dsgo"
	"github.com/zrcoder/dsgo/internal/heap"
)

type Heap[T any] struct {
	cmp  dsgo.Comparator[T]
	data []T
}

func New[T cmp.Ordered](ops ...Option[T]) *Heap[T] {
	return NewWith[T](cmp.Compare[T], ops...)
}

func NewWith[T any](cmp dsgo.Comparator[T], ops ...Option[T]) *Heap[T] {
	h := &Heap[T]{cmp: cmp}
	for _, op := range ops {
		op(h)
	}
	heap.Init(h)
	return h
}

// Push pushes the element value onto the heap.
// The complexity is O(log n) where n is the size of the heap.
func (h *Heap[T]) Push(value T) {
	heap.Push(h, value)
}

// Pop removes and returns the peek element from the heap.
// The complexity is O(log n) where n is the size of the heap.
// Pop is equivalent to Remove(h.data[0]).
func (h *Heap[T]) Pop() (value T, ok bool) {
	if h.Len() == 0 {
		return
	}
	return heap.Pop(h), true
}

// Peek returns the peek value of the heap
// The complexity is O(1)
func (h *Heap[T]) Peek() (value T, ok bool) {
	if h.Len() == 0 {
		return
	}
	return h.data[0], true
}

// implements internal/heap.Interface
func (h *Heap[T]) LenX() int           { return len(h.data) }
func (h *Heap[T]) LessX(i, j int) bool { return h.cmp(h.data[i], h.data[j]) < 0 }
func (h *Heap[T]) SwapX(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap[T]) PushX(x T)           { h.data = append(h.data, x) }
func (h *Heap[T]) PopX() T {
	n := len(h.data)
	x := h.data[n-1]
	var v T
	h.data[n-1] = v // avoid memory leak if T is pointer
	h.data = h.data[:n-1]
	return x
}
