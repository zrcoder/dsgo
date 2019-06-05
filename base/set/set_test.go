package set
 
import "testing"
 
func Test(t *testing.T) {
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
	if set.Exist(300) {
		t.Error("something wrong with func Exist()")
	}
	if !set.Exist(200) {
		t.Error("something wrong with func Exist()")
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
 
func Benchmark(b *testing.B) {
	set := NewWithCapacity(b.N)
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
	for i := 0; i < b.N; i++ {
		set.Delete(i)
	}
}
 
