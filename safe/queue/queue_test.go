package queue

import (
	"testing"
	"sync"
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
	inSequence := make([]int, total)
	outSequencd := make([]int, total)

	wg := sync.WaitGroup{}
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func(item int) {
			queue.Enqueue(item)
			inSequence = append(inSequence, item)
			wg.Done()
		}(i)
	}
	wg.Wait()

	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			item := queue.Dequeue()
			outSequencd = append(outSequencd, item.(int))
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < total; i++ {
		if inSequence[i] != outSequencd[total-i-1] {
			t.Error("failed")
		}
	}
}

func Benchmark(b *testing.B) {
	queue := New()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}
