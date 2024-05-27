package heap

// Remove removes any element from the heap.
// this method is only supported for advanced Heap, it will panics if called on a base heap.
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
// this method is only supported for advanced Heap, it will panics if called on a base heap.
// Changing the value of the element x and then calling Fix is equivalent to,
// but less expensive than, calling Remove(x) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Update(x T) bool {
	if !h.advanced {
		panic("Update is not implemented for base heap")
	}
	if h.cnt[x] == 0 {
		return false
	}
	i := h.idx[x]
	if !h.down(i, len(h.data)) {
		h.up(i)
	}
	return true
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
