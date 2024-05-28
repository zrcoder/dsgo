package ringbuffer

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	buffer := New[int](3)
	if actualValue := buffer.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)

	if actualValue := buffer.Values(); actualValue[0] != 1 || actualValue[1] != 2 || actualValue[2] != 3 {
		t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
	}
	if actualValue := buffer.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := buffer.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := buffer.First(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestFront(t *testing.T) {
	buffer := New[int](3)
	if actualValue, ok := buffer.First(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	buffer.Enqueue(1)
	buffer.Enqueue(2)
	buffer.Enqueue(3)
	if actualValue, ok := buffer.First(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestDequeue(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	buffer := New[int](3)
	assert(buffer.Empty(), true)
	assert(buffer.Empty(), true)
	assert(buffer.Full(), false)
	assert(buffer.Len(), 0)
	buffer.Enqueue(1)
	assert(buffer.Len(), 1)
	buffer.Enqueue(2)
	assert(buffer.Len(), 2)

	buffer.Enqueue(3)
	assert(buffer.Len(), 3)
	assert(buffer.Empty(), false)
	assert(buffer.Full(), true)

	buffer.Dequeue()
	assert(buffer.Len(), 2)

	if actualValue, ok := buffer.First(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	assert(buffer.Len(), 2)

	if actualValue, ok := buffer.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	assert(buffer.Len(), 1)

	if actualValue, ok := buffer.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	assert(buffer.Len(), 0)
	assert(buffer.Empty(), true)
	assert(buffer.Full(), false)

	if actualValue, ok := buffer.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	assert(buffer.Len(), 0)

	assert(buffer.Empty(), true)
	assert(buffer.Full(), false)
	assert(len(buffer.Values()), 0)
}

func TestDequeueFull(t *testing.T) {
	assert := func(actualValue interface{}, expectedValue interface{}) {
		if actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}

	buffer := New[int](2)
	assert(buffer.Empty(), true)
	assert(buffer.Full(), false)
	assert(buffer.Len(), 0)

	buffer.Enqueue(1)
	assert(buffer.Len(), 1)

	buffer.Enqueue(2)
	assert(buffer.Len(), 2)
	assert(buffer.Full(), true)
	if actualValue, ok := buffer.First(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	buffer.Enqueue(3) // overwrites 1
	assert(buffer.Len(), 2)

	if actualValue, ok := buffer.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, expectedValue := buffer.Len(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, ok := buffer.First(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := buffer.Len(), 1; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, ok := buffer.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	assert(buffer.Len(), 0)

	if actualValue, ok := buffer.Dequeue(); actualValue != 0 || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	assert(buffer.Empty(), true)
	assert(buffer.Full(), false)
	assert(len(buffer.Values()), 0)
}

func benchmarkEnqueue(b *testing.B, buffer *Buffer[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			buffer.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, buffer *Buffer[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			buffer.Dequeue()
		}
	}
}

func BenchmarkArrayQueueDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, buffer, size)
}

func BenchmarkArrayQueueDequeue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, buffer, size)
}

func BenchmarkArrayQueueDequeue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, buffer, size)
}

func BenchmarkArrayQueueDequeue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkDequeue(b, buffer, size)
}

func BenchmarkArrayQueueEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	buffer := New[int](3)
	b.StartTimer()
	benchmarkEnqueue(b, buffer, size)
}

func BenchmarkArrayQueueEnqueue1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, buffer, size)
}

func BenchmarkArrayQueueEnqueue10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, buffer, size)
}

func BenchmarkArrayQueueEnqueue100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	buffer := New[int](3)
	for n := 0; n < size; n++ {
		buffer.Enqueue(n)
	}
	b.StartTimer()
	benchmarkEnqueue(b, buffer, size)
}
