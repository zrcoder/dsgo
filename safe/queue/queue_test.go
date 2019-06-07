package queue

import (
	"testing"
	"sync"
	"container/list"
)

func TestBase(t *testing.T) {
	queue := New()
	total := 10
	for i := 0; i < total; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < total; i++ {
		item := queue.Dequeue()
		if i != item {
			t.Errorf("expected %d, but got %d", i, item)
		}
	}
}

func TestSafe(t *testing.T) {
	queue := New()
	const total = 10
		inSequence := list.List{}
		outSequence := list.List{}

	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(item int) {
			queue.Enqueue(item)
			inSequence.PushBack(item)
			wg.Done()
		}(i)
	}
	wg.Wait()

	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			item := queue.Dequeue()
			outSequence.PushBack(item)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < total; i++ {
		inItem := inSequence.Front()
		outItem := outSequence.Front()
		if inItem.Value != outItem.Value {
			t.Error("failed")
		}
		inSequence.Remove(inItem)
		outSequence.Remove(outItem)
	}
}

func Benchmark1(b *testing.B) {
	queue := New()
	wg := sync.WaitGroup{}
	total := b.N
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(item int) {
			queue.Enqueue(item)
			queue.Dequeue()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Benchmark2(b *testing.B) {
	queue := New()
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(item int) {
			queue.Enqueue(item)
			wg.Done()
		}(i)
	}
	wg.Wait()

	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			queue.Dequeue()
			wg.Done()
		}()
	}
	wg.Wait()
}