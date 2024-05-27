package heap

import (
	"math/rand"
	"slices"
	"testing"
)

func TestBinaryHeapPush(t *testing.T) {
	heap := New[int]()

	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	heap.Push(3)
	heap.Push(2)
	heap.Push(1)

	if actualValue, expectedValue := heap.Values(), []int{1, 2, 3}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := heap.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := heap.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := heap.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestBinaryHeapPushBulk(t *testing.T) {
	heap := New[int]()

	heap.Push(15, 20, 3, 1, 2)

	if actualValue, expectedValue := heap.Values(), []int{1, 2, 3, 15, 20}; !slices.Equal(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, ok := heap.Pop(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestBinaryHeapPop(t *testing.T) {
	heap := New[int]()

	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	heap.Pop()

	if actualValue, ok := heap.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := heap.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := heap.Pop(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := heap.Pop(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := heap.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := heap.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func TestBinaryHeapRandom(t *testing.T) {
	heap := New[int]()

	rand.Seed(3)
	for i := 0; i < 10000; i++ {
		r := int(rand.Int31n(30))
		heap.Push(r)
	}

	prev, _ := heap.Pop()
	for !heap.Empty() {
		curr, _ := heap.Pop()
		if prev > curr {
			t.Errorf("Heap property invalidated. prev: %v current: %v", prev, curr)
		}
		prev = curr
	}
}

func benchmarkPush(b *testing.B, heap *Heap[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			heap.Push(n)
		}
	}
}

func benchmarkPop(b *testing.B, heap *Heap[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			heap.Pop()
		}
	}
}

func BenchmarkBinaryHeapPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPop1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPop10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPop100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPop(b, heap, size)
}

func BenchmarkBinaryHeapPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	heap := New[int]()
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkBinaryHeapPush1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkBinaryHeapPush10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, heap, size)
}

func BenchmarkBinaryHeapPush100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	heap := New[int]()
	for n := 0; n < size; n++ {
		heap.Push(n)
	}
	b.StartTimer()
	benchmarkPush(b, heap, size)
}
