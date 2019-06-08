/*
	A stack gives you a LIFO or last-in first-out order.
	You can only Push to add a new element to the top of the stack,
	Pop to remove the element from the top,
	and Peek at the top element without poping it off.
*/
package stack
 
import (
	"sync"
	
	. "github.com/zrcoder/dsGo"
	base "github.com/zrcoder/dsGo/base/stack"
)

type Stack struct {
	lock  sync.RWMutex
	inner *base.Stack
}
 
func New() *Stack {
	return &Stack{inner: base.New()}
}

// Add a new element to the top
func (s *Stack) Push(item Any) {
	s.lock.Lock()
	s.inner.Push(item)
	s.lock.Unlock()
}
 
// Remove the element from the top and returns it
func (s *Stack) Pop() Any {
	s.lock.Lock()
	item := s.inner.Pop()
	s.lock.Unlock()
	return item
}
 
// Returns the element from the top without deletion
func (s *Stack) Peek() Any {
	s.lock.RLock()
	item := s.inner.Peek()
	s.lock.RUnlock()
	return item
}
