package bitset

import (
	"testing"
	"sync"
)

func Test(t *testing.T) {
	const total = 2019
	bs := NewWithSize(total)
	for i := uint(0); i < total; i += 2 {
		bs.Set(i)
	}
	for i := uint(0); i < total; i++ {
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
	total := uint(b.N)
	bs := NewWithSize(uint(total))
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := uint(0); i < total; i ++ {
		go bs.Set(i)
		go bs.Get(i)
		wg.Done()
	}
	wg.Wait()
}
