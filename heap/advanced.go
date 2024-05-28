package heap

// Remove removes value (any element) from the heap.
// This method is only supported for advanced Heap, it will panics if called on a base heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Remove(value T) bool {
	if !h.advanced {
		panic("Remove is not implemented for base heap")
	}
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
	if !h.advanced {
		panic("Update is not implemented for base heap")
	}
	if h.cnt[value] == 0 {
		return false
	}
	i := h.idx[value]
	return h.updateIndex(i)
}

func (h *Heap[T]) removeIndex(index int) bool {
	if index < 0 || index >= len(h.data) {
		return false
	}
	do := func() {
		last := len(h.data) - 1
		if last != index {
			h.swap(index, last)
			if !h.down(index, last) {
				h.up(index)
			}
		}
		var zero T
		h.data[last] = zero // avoid memory leak if T is pointer
		h.data = h.data[:last]
	}
	if !h.advanced {
		do()
		return true
	}
	value := h.data[index]
	if h.cnt[value] == 1 {
		do()
		delete(h.idx, value)
	}
	h.size--
	h.cnt[value]--
	return true
}

func (h *Heap[T]) updateIndex(index int) bool {
	if !h.down(index, len(h.data)) {
		h.up(index)
	}
	return true
}
