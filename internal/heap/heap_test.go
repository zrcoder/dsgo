package heap

import (
	"cmp"
	"math/rand"
	"testing"
)

type Heap[T cmp.Ordered] struct {
	s []T
}

func (h *Heap[T]) LessX(i, j int) bool {
	return h.s[i] < h.s[j]
}

func (h *Heap[T]) SwapX(i, j int) {
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

func (h *Heap[T]) LenX() int {
	return len(h.s)
}

func (h *Heap[T]) PopX() (v T) {
	n := len(h.s)
	x := h.s[n-1]
	h.s = h.s[:n-1]
	return x
}

func (h *Heap[T]) PushX(v T) {
	h.s = append(h.s, v)
}

func (h *Heap[T]) verify(t *testing.T, i int) {
	t.Helper()
	n := h.LenX()
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h.LessX(j1, i) {
			t.Errorf("heap invariant invalidated [%d] = %v > [%d] = %v", i, h.s[i], j1, h.s[j1])
			return
		}
		h.verify(t, j1)
	}
	if j2 < n {
		if h.LessX(j2, i) {
			t.Errorf("heap invariant invalidated [%d] = %v > [%d] = %v", i, h.s[i], j1, h.s[j2])
			return
		}
		h.verify(t, j2)
	}
}

func TestInit0(t *testing.T) {
	h := new(Heap[int])
	for i := 20; i > 0; i-- {
		h.PushX(0) // all elements are the same
	}
	Init(h)
	h.verify(t, 0)

	for i := 1; h.LenX() > 0; i++ {
		x := Pop(h)
		h.verify(t, 0)
		if x != 0 {
			t.Errorf("%d.th pop got %d; want %d", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	h := new(Heap[int])
	for i := 20; i > 0; i-- {
		h.PushX(i) // all elements are different
	}
	Init(h)
	h.verify(t, 0)

	for i := 1; h.LenX() > 0; i++ {
		x := Pop(h)
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func Test(t *testing.T) {
	h := new(Heap[int])
	h.verify(t, 0)

	for i := 20; i > 10; i-- {
		h.PushX(i)
	}
	Init(h)
	h.verify(t, 0)

	for i := 10; i > 0; i-- {
		Push(h, i)
		h.verify(t, 0)
	}

	for i := 1; h.LenX() > 0; i++ {
		x := Pop(h)
		if i < 20 {
			Push(h, 20+i)
		}
		h.verify(t, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}

func TestRemove0(t *testing.T) {
	h := new(Heap[int])
	for i := 0; i < 10; i++ {
		h.PushX(i)
	}
	h.verify(t, 0)

	for h.LenX() > 0 {
		i := h.LenX() - 1
		x := RemoveIndex(h, i)
		if x != i {
			t.Errorf("Remove(%d) got %d; want %d", i, x, i)
		}
		h.verify(t, 0)
	}
}

func TestRemove1(t *testing.T) {
	h := new(Heap[int])
	for i := 0; i < 10; i++ {
		h.PushX(i)
	}
	h.verify(t, 0)

	for i := 0; h.LenX() > 0; i++ {
		x := RemoveIndex(h, 0)
		if x != i {
			t.Errorf("Remove(0) got %d; want %d", x, i)
		}
		h.verify(t, 0)
	}
}

func TestRemove2(t *testing.T) {
	N := 10

	h := new(Heap[int])
	for i := 0; i < N; i++ {
		h.PushX(i)
	}
	h.verify(t, 0)

	m := make(map[int]bool)
	for h.LenX() > 0 {
		m[RemoveIndex(h, (h.LenX()-1)/2)] = true
		h.verify(t, 0)
	}

	if len(m) != N {
		t.Errorf("len(m) = %d; want %d", len(m), N)
	}
	for i := 0; i < len(m); i++ {
		if !m[i] {
			t.Errorf("m[%d] doesn't exist", i)
		}
	}
}

func BenchmarkDup(b *testing.B) {
	const n = 10000
	h := &Heap[int]{s: make([]int, 0, n)}
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			Push(h, 0) // all elements are the same
		}
		for h.LenX() > 0 {
			Pop(h)
		}
	}
}

func TestFix(t *testing.T) {
	h := new(Heap[int])
	h.verify(t, 0)

	for i := 200; i > 0; i -= 10 {
		Push(h, i)
	}
	h.verify(t, 0)

	if h.s[0] != 10 {
		t.Fatalf("Expected head to be 10, was %d", h.s[0])
	}
	h.s[0] = 210
	FixIndex(h, 0)
	h.verify(t, 0)

	for i := 100; i > 0; i-- {
		elem := rand.Intn(h.LenX())
		if i&1 == 0 {
			h.s[elem] *= 2
		} else {
			h.s[elem] /= 2
		}
		FixIndex(h, elem)
		h.verify(t, 0)
	}
}
