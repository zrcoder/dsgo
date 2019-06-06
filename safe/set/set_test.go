package set
 
import (
	"testing"
	"sync"
)
 
func TestBase(t *testing.T) {
	set := New()
	for i := 0; i < 10; i++ {
		set.Add(i)
	}
	if set.Count() != 10 {
		t.Error("something wrong with func Count()")
	}
	set.Add(100)
	set.Add(100)
	set.Add(200)
	if set.Has(300) {
		t.Error("something wrong with func Has()")
	}
	if !set.Has(200) {
		t.Error("something wrong with func Has()")
	}
	set.Delete(100)
	set.Delete(555)
	if set.Count() != 11 {
		t.Error("something wrong with func Count()")
	}
	if len(set.AllItems()) != set.Count() {
		t.Error("len(set.AllItems()) != set.Count()")
	}
}
 
func TestSafe(t *testing.T) {
	set := New()
	const max = 20
	const itemToDel = 5
	wg := sync.WaitGroup{}
	wg.Add(max * 2)
 
	for i := 1; i <= max; i++ {
		go func(i int) {
			set.Add(i)
			wg.Done()
		}(i)
		go func() {
			set.Delete(itemToDel)
			wg.Done()
		}()
	}
 
	wg.Wait()
 
	if set.Count() != max-1 {
		t.Errorf("length is %d, expected %d\n", set.Count(), max-1)
	}
	if set.Has(itemToDel) {
		t.Errorf("the item %d exists, expected not exsit", itemToDel)
	}
	t.Log(set.AllItems())
}
 
func Benchmark(b *testing.B) {
	set := NewWithCapacity(b.N)
 
	for i := 0; i < b.N; i++ {
		go set.Add(i)
		go set.Has(i)
	}
}
