package bidmap

import (
	"strings"
	"testing"
)

func TestMapPut(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") // overwrite

	if actualValue := m.Len(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}

	checkKeys(t, m, []int{1, 2, 3, 4, 5, 6, 7})
	checkValues(t, m, []string{"a", "b", "c", "d", "e", "f", "g"})

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test[0].(int))
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}
}

func checkKeys[K int, V string](t *testing.T, m *Map[K, V], keys []K) {
	for _, key := range keys {
		if _, ok := m.Get(key); !ok {
			t.Errorf("%d is expected in the map, but not", key)
		}
	}
}

func checkValues[K int, V string](t *testing.T, m *Map[K, V], values []V) {
	for _, value := range values {
		if _, ok := m.GetKey(value); !ok {
			t.Errorf("%s is expected in the map, but not", value)
		}
	}
}

func TestMapRemove(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") // overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	checkKeys(t, m, []int{1, 2, 3, 4})
	checkValues(t, m, []string{"a", "b", "c", "d"})

	if actualValue := m.Len(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	tests2 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "", false},
		{6, "", false},
		{7, "", false},
		{8, "", false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := m.Get(test[0].(int))
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	if len(m.Keys()) > 0 {
		t.Error("expect keys length 0")
	}
	if len(m.Values()) > 0 {
		t.Error("expect values length 0")
	}
	if actualValue := m.Len(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapGetKey(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") // overwrite

	// key,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{0, "x", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.GetKey(test[1].(string))
		if actualValue != test[0] || actualFound != test[2] {
			t.Errorf("Got %v expected %v for value %s", actualValue, test[0], test[1])
		}
	}
}

func TestMapString(t *testing.T) {
	c := New[string, int]()
	c.Put("a", 1)
	if !strings.HasPrefix(c.String(), "BidMap") {
		t.Errorf("String should start with container name")
	}
}

func benchmarkGet(b *testing.B, m *Map[int, int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map[int, int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, n)
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map[int, int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkHashBidiMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkHashBidiMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, int]()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkHashBidiMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashBidiMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashBidiMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkHashBidiMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, int]()
	for n := 0; n < size; n++ {
		m.Put(n, n)
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
