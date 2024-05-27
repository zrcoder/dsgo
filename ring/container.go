package ring

import "github.com/zrcoder/dsgo"

var _ dsgo.Container[int] = (*Ring[int])(nil)

// Len computes the number of elements in ring r.
// It executes in time proportional to the number of elements.
func (r *Ring[T]) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func (r *Ring[T]) Empty() bool {
	return r.Len() == 0
}

func (r *Ring[T]) Values() []T {
	res := make([]T, r.Len())
	p := r
	for i := range res {
		res[i] = p.Value
		p = p.next
	}
	return res
}

func (r *Ring[T]) Clear() {
	r.init()
}
