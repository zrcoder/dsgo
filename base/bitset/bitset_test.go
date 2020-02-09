package bitset

import "testing"

func Test(t *testing.T) {
	const total = 2019
	bs := NewWithSize(total)
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

func Benchmark(b *testing.B) {
	bs := NewWithSize(b.N)
	for i := 0; i < b.N; i++ {
		bs.Set(i)
	}
	for i := 0; i < b.N; i++ {
		bs.Get(i)
	}
}
