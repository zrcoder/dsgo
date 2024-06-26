package heap

import "github.com/zrcoder/dsgo"

type Option[T any] func(h *Heap[T])

func WithComparator[T any](cmp dsgo.Comparator[T]) Option[T] {
	return func(h *Heap[T]) {
		h.cmp = cmp
	}
}

func WithCapacity[T any](capacity int) Option[T] {
	return func(h *Heap[T]) {
		if capacity <= len(h.data) {
			h.data = h.data[:capacity:capacity]
			return
		}
		if capacity <= cap(h.data) {
			h.data = h.data[:len(h.data):capacity]
			return
		}
		data := make([]T, len(h.data), capacity)
		copy(data, h.data)
		h.data = data
	}
}

func WithData[T any, S ~[]T](data S) Option[T] {
	return func(h *Heap[T]) {
		if cap(h.data) < len(data) {
			h.data = make([]T, len(data))
		} else {
			h.data = h.data[:len(data)]
		}
		copy(h.data, data)
	}
}
