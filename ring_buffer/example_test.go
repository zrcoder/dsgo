package ring_buffer

import "fmt"

func Example() {
	buf := New[int](4)

	buf.Enqueue(123)
	buf.Enqueue(456)
	buf.Enqueue(789)
	buf.Enqueue(666)

	fmt.Println(buf.Dequeue())
	fmt.Println(buf.Dequeue())
	fmt.Println(buf.Dequeue())

	buf.Enqueue(333)
	buf.Enqueue(555)

	for !buf.Empty() {
		fmt.Println(buf.Dequeue())
	}

	// output:
	// 123 true
	// 456 true
	// 789 true
	// 666 true
	// 333 true
	// 555 true
}
