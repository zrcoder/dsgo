/*
A queue gives you a FIFO or first-in firs-out order.
The element you inserted first is also the first one to come out again.
*/
package queue

import (
	"github.com/zrcoder/dsgo/list"
)

// Queue is a FIFO (first-in first-out) list
type Queue[T any] struct {
	list *list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{list: list.New[T]()}
}

// Enqueue puts a given item at the back of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.list.PushBack(item)
}

// Dequeue removes the first item from the queue and returns it
func (q *Queue[T]) Dequeue() (value T, ok bool) {
	if q.Empty() {
		return
	}
	return q.list.Remove(q.list.Front()), true
}

// Front returns the first item of the queue
func (q *Queue[T]) Front() (value T, ok bool) {
	if q.Empty() {
		return
	}
	return q.list.Front().Value, true
}
