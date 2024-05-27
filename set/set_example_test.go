package set

import (
	"fmt"
)

func Example_range() {
	const tolal = 10
	set := NewWithCapacity[int](tolal)
	for i := 1; i <= tolal; i++ {
		set.Add(i)
	}
	sum := 0
	set.Range(func(item int) bool {
		sum += item
		return false
	})
	fmt.Println(sum)

	// OutPut:
	// 55
}
