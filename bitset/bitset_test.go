package bitset

import "testing"

func Test(t *testing.T) {
	const total = 2019
	bs := New(total)
	for i := 0; i < total; i += 2 {
		bs.Set(i)
	}
	for i := 0; i < total; i++ {
		seted := bs.Get(i)
		switch {
		case i%2 == 0 && !seted:
			t.Error("failed")
		case i%2 == 1 && seted:
			t.Error("failed")
		}
	}
	bs.Unset(8)
	if bs.Get(8) {
		t.Error("failed")
	}
}

func gen(size int, idx ...int) BitSet {
	res := New(size)
	for _, i := range idx {
		res.Set(i)
	}
	return res
}

func TestIntersection(t *testing.T) {
	a := gen(10, 0, 7)
	b := gen(20, 1, 7, 18)
	x := Intersection(a, b)
	if len(x) != len(a) {
		t.Errorf("expected len %d, got %d.", len(a), len(x))
	}
	for i := range x {
		val := x.Get(i)
		if i == 7 {
			if !val {
				t.Errorf("expected true, got false on bit %d", i)
			}
		} else if val {
			t.Errorf("expected false, got true on bit %d", i)
		}
	}
}

func TestUnion(t *testing.T) {
	a := gen(10, 0, 7)
	b := gen(20, 1, 7, 18)
	x := Union(a, b)
	if len(x) != len(b) {
		t.Errorf("expected len %d, got %d.", len(b), len(x))
	}
	for i := range x {
		val := x.Get(i)
		switch i {
		case 0, 1, 7, 18:
			if !val {
				t.Errorf("expected true, got false on bit %d", i)
			}
		default:
			if val {
				t.Errorf("expected false, got true on bit %d", i)
			}
		}
	}
}

func TestDifference(t *testing.T) {
	a := gen(10, 0, 7)
	b := gen(20, 1, 7, 18)
	x := Difference(a, b)
	if len(x) != len(b) {
		t.Errorf("expected len %d, got %d.", len(b), len(x))
	}
	for i := range x {
		val := x.Get(i)
		switch i {
		case 0, 1, 18:
			if !val {
				t.Errorf("expected true, got false on bit %d", i)
			}
		default:
			if val {
				t.Errorf("expected false, got true on bit %d", i)
			}
		}
	}
}

func Benchmark(b *testing.B) {
	bs := New(b.N)
	for i := 0; i < b.N; i++ {
		bs.Set(i)
	}
	for i := 0; i < b.N; i++ {
		bs.Get(i)
	}
}
