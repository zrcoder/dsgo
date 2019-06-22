package set

import (
	"fmt"
	"sync"

	. "github.com/zrcoder/dsGo"
)

func ExampleRange() {
	const tolal = 10
	set := NewWithCapacity(tolal)
	wg := sync.WaitGroup{}
	wg.Add(tolal)
	for i := 1; i <= tolal; i++ {
		go func(item int) {
			set.Add(item)
			wg.Done()
		}(i)
	}
	wg.Wait()

	sum := 0
	set.Range(func(item Any) bool {
		switch v := item.(type) {
		case int:
			sum += v
			return false
		default:
			return true
		}
	})

	fmt.Println(sum)

	// OutPut:
	// 55
}
