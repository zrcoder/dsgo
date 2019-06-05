package stack
 
import (
	"testing"
	"sync"
)
 
func Test(t *testing.T) {
	stack := New()
	const total = 10
	for i := 0; i < total; i++ {
		stack.Push(i)
	}
	for i := total - 1; i > -1; i-- {
		item := stack.Pop()
		if item != i {
			t.Error("failed")
		}
	}
}
 
func TestSafe(t *testing.T) {
	stack := New()
	const total = 10
	pushed := make([]int, total)
	poped := make([]int, total)
 
	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(i int) {
			stack.Push(i)
			pushed = append(pushed, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
 
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			item := stack.Pop()
			poped = append(poped, item.(int))
			wg.Done()
		}()
	}
	wg.Wait()
 
	for i:=0; i< total; i++ {
		if pushed[i] != poped[total-i-1] {
			t.Error("failed")
		}
	}
}
 
func Benchmark(b *testing.B) {
	stack := New()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	for i := b.N - 1; i > -1; i-- {
		stack.Pop()
	}
}
