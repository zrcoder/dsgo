package heap

import (
	"cmp"
	"fmt"

	"github.com/zrcoder/dsgo"
)

func Example_ints() {
	h := New[int]()
	h.Push(2)
	h.Push(1)
	h.Push(5)
	h.Push(3)
	peek, _ := h.Peek()
	fmt.Printf("minimum: %d\n", peek)
	for h.Len() > 0 {
		cur, _ := h.Pop()
		fmt.Printf("%d ", cur)
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func Example_withComparator() {
	h := New(WithComparator(dsgo.Reverse(cmp.Compare[int])))
	h.Push(2)
	h.Push(1)
	h.Push(5)
	h.Push(3)
	peek, _ := h.Peek()
	fmt.Printf("maximum: %d\n", peek)
	for h.Len() > 0 {
		cur, _ := h.Pop()
		fmt.Printf("%d ", cur)
	}
	// Output:
	// maximum: 5
	// 5 3 2 1
}

type Item struct {
	Name     string
	Priority int
}

func Example_custom() {
	h := NewWith(func(a, b Item) int {
		return b.Priority - a.Priority
	})
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	for name, priority := range items {
		h.Push(Item{
			Name:     name,
			Priority: priority,
		})
	}
	item := Item{
		Name:     "orange",
		Priority: 1,
	}
	h.Push(item)
	item.Priority = 5
	h.Push(item)
	for h.Len() > 0 {
		item, _ := h.Pop()
		fmt.Printf("%.2d:%s ", item.Priority, item.Name)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple 01:orange
}

func Example_withData() {
	nums := []int{6, 8, 5, 9, 3}
	h := New(WithData(nums), WithCapacity[int](len(nums)+1))
	h.Push(1)
	for h.Len() > 0 {
		cur, _ := h.Pop()
		fmt.Print(cur)
		fmt.Print(",")
	}
	// Output:
	// 1,3,5,6,8,9,
}
