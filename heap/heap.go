package heap

import (
	"cmp"

	"github.com/zrcoder/dsgo"
)

type Heap[T comparable] struct {
	cmp      dsgo.Comparator[T]
	data     []T
	idx      map[T]int
	cnt      map[T]int
	size     int
	advanced bool
}

func New[T cmp.Ordered](ops ...Option[T]) *Heap[T] {
	return NewWith[T](dsgo.OrderedComparator[T](), ops...)
}

func NewWith[T comparable](cmp dsgo.Comparator[T], ops ...Option[T]) *Heap[T] {
	res := &Heap[T]{
		cmp: cmp,
	}
	for _, op := range ops {
		op(res)
	}
	if res.advanced {
		res.size = len(res.data)
		res.idx = make(map[T]int)
		res.cnt = make(map[T]int)
	}
	res.build()
	return res
}

// build will build the heap by the given cmp.
// The complexity is O(n) where n is the size of the heap.
func (h *Heap[T]) build() {
	if h.advanced {
		data := make([]T, 0, len(h.data))
		for _, v := range h.data {
			h.cnt[v]++
			if h.cnt[v] == 1 {
				data = append(data, v)
			}
		}
		h.data = data
	}
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	if h.advanced {
		for i, v := range h.data {
			h.idx[v] = i
		}
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
	if h.advanced {
		if h.cnt[value] == 0 {
			h.idx[value] = n
			h.data = append(h.data, value)
			h.up(n)
		}
		h.cnt[value]++
		h.size++
	} else {
		h.data = append(h.data, value)
		h.up(n)
	}
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

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	if h.advanced {
		h.idx[h.data[i]] = i
		h.idx[h.data[j]] = j
	}
}

func (h *Heap[T]) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || h.cmp(h.data[i], h.data[parent]) >= 0 {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

func (h *Heap[T]) down(i, n int) bool {
	cur := i
	for {
		child := 2*cur + 1 // left
		if child >= n || child < 0 {
			break
		}
		right := child + 1
		if right < n && h.cmp(h.data[right], h.data[child]) < 0 {
			child = right
		}
		if h.cmp(h.data[child], h.data[cur]) >= 0 {
			break
		}
		h.swap(cur, child)
		cur = child
	}
	return cur > i
}
