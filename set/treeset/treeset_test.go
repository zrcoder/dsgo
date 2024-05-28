package treeset

import (
	"slices"
	"strings"
	"testing"
)

func TestSetNew(t *testing.T) {
	set := New[int](2, 1)
	if actualValue := set.Len(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	values := set.Values()
	if actualValue := values[0]; actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue := values[1]; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}

func TestSetAdd(t *testing.T) {
	set := New[int]()
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	values := set.Values()
	expected := []int{1, 2, 3}
	if !slices.Equal(values, expected) {
		t.Errorf("Got %v expected %v", values, expected)
	}
}

func TestSetContains(t *testing.T) {
	set := New[int]()
	set.Add(3, 1, 2)
	if actualValue := set.Contains(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestSetRemove(t *testing.T) {
	set := New[int]()
	set.Add(3, 1, 2)
	set.Remove()
	if actualValue := set.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	set.Remove(1)
	if actualValue := set.Len(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	set.Remove(3)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if actualValue := set.Len(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSetChaining(t *testing.T) {
	set := New[string]()
	set.Add("c", "a", "b")
}

func TestSetString(t *testing.T) {
	c := New[int]()
	c.Add(1)
	if !strings.HasPrefix(c.String(), "TreeSet") {
		t.Errorf("String should start with container name")
	}
}

func TestSetIntersection(t *testing.T) {
	set := New[string]()
	another := New[string]()

	intersection := set.Intersection(another)
	if actualValue, expectedValue := intersection.Len(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	intersection = set.Intersection(another)

	if actualValue, expectedValue := intersection.Len(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := intersection.Contains("c", "d"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetUnion(t *testing.T) {
	set := New[string]()
	another := New[string]()

	union := set.Union(another)
	if actualValue, expectedValue := union.Len(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	union = set.Union(another)

	if actualValue, expectedValue := union.Len(), 6; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := union.Contains("a", "b", "c", "d", "e", "f"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestSetDifference(t *testing.T) {
	set := New[string]()
	another := New[string]()

	difference := set.Difference(another)
	if actualValue, expectedValue := difference.Len(), 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	set.Add("a", "b", "c", "d")
	another.Add("c", "d", "e", "f")

	difference = set.Difference(another)

	if actualValue, expectedValue := difference.Len(), 2; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := difference.Contains("a", "b"); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func benchmarkContains(b *testing.B, set *Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *Set[int], size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkTreeSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New[int]()
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTreeSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTreeSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTreeSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New[int]()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}
