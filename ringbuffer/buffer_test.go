package ringbuffer

import (
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := New[int](3)
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
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := New[int](3)
	if actualValue, ok := queue.Peek(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueueDequeue(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	queue := New[int](3)
	assert(queue.Empty(), true)
	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(queue.Len(), 0)
	queue.Enqueue(1)
	assert(queue.Len(), 1)
	queue.Enqueue(2)
	assert(queue.Len(), 2)

	queue.Enqueue(3)
	assert(queue.Len(), 3)
	assert(queue.Empty(), false)
	assert(queue.Full(), true)

	queue.Dequeue()
	assert(queue.Len(), 2)

	if actualValue, ok := queue.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	assert(queue.Len(), 2)

	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	assert(queue.Len(), 1)

	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	assert(queue.Len(), 0)
	assert(queue.Empty(), true)
	assert(queue.Full(), false)

	if actualValue, ok := queue.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	assert(queue.Len(), 0)

	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(len(queue.Values()), 0)
}

func TestQueueDequeueFull(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	queue := New[int](2)
	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(queue.Len(), 0)

	queue.Enqueue(1)
	assert(queue.Len(), 1)

	queue.Enqueue(2)
	assert(queue.Len(), 2)
	assert(queue.Full(), true)
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	queue.Enqueue(3) // overwrites 1
	assert(queue.Len(), 2)

	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, expectedValue := queue.Len(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, ok := queue.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := queue.Len(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	assert(queue.Len(), 0)

	if actualValue, ok := queue.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	assert(queue.Empty(), true)
	assert(queue.Full(), false)
	assert(len(queue.Values()), 0)
}

func benchmarkEnqueue(b *testing.B, queue *Buffer[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, queue *Buffer[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			queue.Dequeue()
		}
	}
}

func BenchmarkArrayQueueDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueDequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	queue := New[int](3)
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}

func BenchmarkArrayQueueEnqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	queue := New[int](3)
	for n := 0; n < size; n++ {
		queue.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, queue, size)
}
