package queue

import (
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := New[int]()
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if actualValue := queue.Values(); actualValue[0] != 1 || actualValue[1] != 2 || actualValue[2] != 3 {
		t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
	}
	if actualValue := queue.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := queue.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueueFront(t *testing.T) {
	queue := New[int]()
	if actualValue, ok := queue.Front(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if actualValue, ok := queue.Front(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueueDequeue(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	if actualValue, ok := queue.Front(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := queue.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}

func benchmarkEnqueue(b *testing.B, queue *Queue[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, queue *Queue[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Dequeue()
		}
	}
}

func BenchmarkArrayQueueDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New[int]()
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New[int]()
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}
