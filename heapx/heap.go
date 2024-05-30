// Package heapx is an advanced heap, which can remove or update any item in O(nlogn) time.
// It uses hash maps to memory inner index and count for any item.
package heapx

import (
	"cmp"

	"github.com/zrcoder/dsgo"
	"github.com/zrcoder/dsgo/internal/heap"
)

type Heap[T comparable] struct {
	cmp  dsgo.Comparator[T]
	data []T
	idx  map[T]int
	cnt  map[T]int
	size int
}

func New[T cmp.Ordered](ops ...Option[T]) *Heap[T] {
	return NewWith[T](cmp.Compare[T], ops...)
}

func NewWith[T comparable](cmp dsgo.Comparator[T], ops ...Option[T]) *Heap[T] {
	h := &Heap[T]{cmp: cmp}
	for _, op := range ops {
		op(h)
	}
	h.size = len(h.data)
	h.idx = make(map[T]int)
	h.cnt = make(map[T]int)
	h.build()
	return h
}

// build will build the heap by the given cmp.
// The complexity is O(n) where n is the size of the heap.
func (h *Heap[T]) build() {
	data := make([]T, 0, len(h.data))
	for _, v := range h.data {
		h.cnt[v]++
		if h.cnt[v] == 1 {
			data = append(data, v)
		}
	}
	h.data = data
	heap.Init(h)
	for i, v := range h.data {
		h.idx[v] = i
	}
}

// Push pushes the elements values onto the heap.
// The complexity is O(mlog n) where m is the size of values and n is the size of the heap.
func (h *Heap[T]) Push(values ...T) {
	for _, value := range values {
		h.push(value)
	}
}

// push pushes the element x onto the heap.
// The complexity is O(log n) where n is the size of the heap.
func (h *Heap[T]) push(value T) {
	n := len(h.data)
	if h.cnt[value] == 0 {
		h.idx[value] = n
		heap.Push(h, value)
	}
	h.cnt[value]++
	h.size++
}

// Pop removes and returns the peek element from the heap.
// The complexity is O(log n) where n is the size of the heap.
// Pop is equivalent to Remove(h.data[0]).
func (h *Heap[T]) Pop() (value T, ok bool) {
	if h.Len() == 0 {
		return
	}
	value = h.data[0]
	h.removeIndex(0)
	return value, true
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

func (h *Heap[T]) SwapX(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.idx[h.data[i]] = i
	h.idx[h.data[j]] = j
}

func (h *Heap[T]) PushX(x T) { h.data = append(h.data, x) }

func (h *Heap[T]) PopX() T {
	n := len(h.data)
	x := h.data[n-1]
	var v T
	h.data[n-1] = v // avoid memory leak if T is pointer
	h.data = h.data[:n-1]
	return x
}
