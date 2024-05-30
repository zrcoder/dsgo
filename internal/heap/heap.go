// Package heap is a generics version of "container/heap" in the standard library
package heap

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// [Init] has been called or if the data is empty or sorted):
//
//	!h.LessX(j, i) for 0 <= i < h.LenX() and 2*i+1 <= j <= 2*i+2 and j < h.LenX()
//
// Note that [PushX] and [PopX] in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use [heap.Push] and [heap.Pop].
type Interface[T any] interface {
	LenX() int
	LessX(int, int) bool
	SwapX(int, int)
	PushX(T)
	PopX() T
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init[T any](h Interface[T]) {
	// heapify
	n := h.LenX()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push[T any](h Interface[T], x T) {
	h.PushX(x)
	up(h, h.LenX()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to [Remove](h, 0).
func Pop[T any](h Interface[T]) T {
	n := h.LenX() - 1
	h.SwapX(0, n)
	down(h, 0, n)
	return h.PopX()
}

// RemoveIndex removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func RemoveIndex[T any](h Interface[T], i int) T {
	n := h.LenX() - 1
	if n != i {
		h.SwapX(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.PopX()
}

// FixIndex re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling [Remove](h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func FixIndex[T any](h Interface[T], i int) {
	if !down(h, i, h.LenX()) {
		up(h, i)
	}
}

func up[T any](h Interface[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.LessX(j, i) {
			break
		}
		h.SwapX(i, j)
		j = i
	}
}

func down[T any](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.LessX(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.LessX(j, i) {
			break
		}
		h.SwapX(i, j)
		i = j
	}
	return i > i0
}
