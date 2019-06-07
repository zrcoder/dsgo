package stack
 
import (
	"testing"
	"sync"
	"container/list"
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
	pushed := list.New()
	poped := list.New()
 
	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(i int) {
			stack.Push(i)
			pushed.PushBack(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
 
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			item := stack.Pop()
			poped.PushBack(item)
			wg.Done()
		}()
	}
	wg.Wait()
 
	for i:=0; i< total; i++ {
		pushedItem := pushed.Front()
		popedItem := poped.Back()
		if pushedItem.Value != popedItem.Value {
			t.Error("failed")
		}
	}
}
 
func Benchmark(b *testing.B) {
	stack := New()
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go stack.Push(i)
		go stack.Pop()
		wg.Done()
	}
	wg.Wait()
}
