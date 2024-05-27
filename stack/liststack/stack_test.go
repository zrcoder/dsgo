package liststack

import "testing"

func Test(t *testing.T) {
	stack := New[int]()
	const total = 10
	for i := 0; i < total; i++ {
		stack.Push(i)
	}
	for i := total - 1; i > -1; i-- {
		item, _ := stack.Pop()
		if item != i {
			t.Error("failed")
		}
	}
}

func Benchmark(b *testing.B) {
	stack := New[int]()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	for i := b.N - 1; i > -1; i-- {
		stack.Pop()
	}
}
