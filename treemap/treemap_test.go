package treemap

import (
	"slices"
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
	keys, expectedKeys := m.Keys(), []int{1, 2, 3, 4, 5, 6, 7}
	if !slices.Equal(keys, expectedKeys) {
		t.Errorf("Got keys %v expected %v", keys, expectedKeys)
	}
	vals, expectedVals := m.Values(), []string{"a", "b", "c", "d", "e", "f", "g"}
	if !slices.Equal(vals, expectedVals) {
		t.Errorf("Got values %v expected %v", vals, expectedVals)
	}

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

func TestMapMin(t *testing.T) {
	m := New[int, string]()

	if k, v, ok := m.Min(); k != 0 || v != "" || ok {
		t.Errorf("Got %v->%v->%v expected %v->%v-%v", k, v, ok, 0, "", false)
	}

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") // overwrite

	actualKey, actualValue, actualOk := m.Min()
	expectedKey, expectedValue, expectedOk := 1, "a", true
	if actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualOk != expectedOk {
		t.Errorf("Got %v expected %v", actualOk, expectedOk)
	}
}

func TestMapMax(t *testing.T) {
	m := New[int, string]()

	if k, v, ok := m.Max(); k != 0 || v != "" || ok {
		t.Errorf("Got %v->%v->%v expected %v->%v-%v", k, v, ok, 0, "", false)
	}

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") // overwrite

	actualKey, actualValue, actualOk := m.Max()
	expectedKey, expectedValue, expectedOk := 7, "g", true
	if actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualOk != expectedOk {
		t.Errorf("Got %v expected %v", actualOk, expectedOk)
	}
}

func TestMapClear(t *testing.T) {
	m := New[int, string]()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	if actualValue, expectedValue := m.Len(), 4; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	m.Clear()
	if actualValue, expectedValue := m.Len(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
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

	keys, expectedKeys := m.Keys(), []int{1, 2, 3, 4}
	if !slices.Equal(keys, expectedKeys) {
		t.Errorf("Got keys %v, expected %v", keys, expectedKeys)
	}
	vals, expectedVals := m.Values(), []string{"a", "b", "c", "d"}
	if !slices.Equal(vals, expectedVals) {
		t.Errorf("Got values %v, expected %v", vals, expectedVals)
	}

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
	keys = m.Keys()
	if len(keys) > 0 {
		t.Errorf("Got keys %v, expected empty", keys)
	}
	vals = m.Values()
	if len(vals) > 0 {
		t.Errorf("Got values %v, expected empty", vals)
	}

	if actualValue := m.Len(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestMapFloor(t *testing.T) {
	m := New[int, string]()
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(1, "a")

	// key,expectedKey,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{-1, 0, "", false},
		{0, 0, "", false},
		{1, 1, "a", true},
		{2, 1, "a", true},
		{3, 3, "c", true},
		{4, 3, "c", true},
		{7, 7, "g", true},
		{8, 7, "g", true},
	}

	for _, test := range tests1 {
		// retrievals
		actualKey, actualValue, actualOk := m.Floor(test[0].(int))
		if actualKey != test[1] || actualValue != test[2] || actualOk != test[3] {
			t.Errorf("Got %v, %v, %v, expected %v, %v, %v", actualKey, actualValue, actualOk, test[1], test[2], test[3])
		}
	}
}

func TestMapCeiling(t *testing.T) {
	m := New[int, string]()
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(1, "a")

	// key,expectedKey,expectedValue,expectedFound
	tests1 := [][]interface{}{
		{-1, 1, "a", true},
		{0, 1, "a", true},
		{1, 1, "a", true},
		{2, 3, "c", true},
		{3, 3, "c", true},
		{4, 7, "g", true},
		{7, 7, "g", true},
		{8, 0, "", false},
	}

	for _, test := range tests1 {
		// retrievals
		actualKey, actualValue, actualOk := m.Ceiling(test[0].(int))
		if actualKey != test[1] || actualValue != test[2] || actualOk != test[3] {
			t.Errorf("Got %v, %v, %v, expected %v, %v, %v", actualKey, actualValue, actualOk, test[1], test[2], test[3])
		}
	}
}

func TestMapString(t *testing.T) {
	c := New[string, int]()
	c.Put("a", 1)
	if !strings.HasPrefix(c.String(), "TreeMap") {
		t.Errorf("String should start with container name")
	}
}

// noinspection GoBoolExpressions
func assertSerialization(m *Map[string, string], txt string, t *testing.T) {
	if actualValue := m.Keys(); false ||
		actualValue[0] != "a" ||
		actualValue[1] != "b" ||
		actualValue[2] != "c" ||
		actualValue[3] != "d" ||
		actualValue[4] != "e" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[a,b,c,d,e]")
	}
	if actualValue := m.Values(); false ||
		actualValue[0] != "1" ||
		actualValue[1] != "2" ||
		actualValue[2] != "3" ||
		actualValue[3] != "4" ||
		actualValue[4] != "5" {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, "[1,2,3,4,5]")
	}
	if actualValue, expectedValue := m.Len(), 5; actualValue != expectedValue {
		t.Errorf("[%s] Got %v expected %v", txt, actualValue, expectedValue)
	}
}

func benchmarkGet(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map[int, struct{}], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkTreeMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := New[int, struct{}]()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
