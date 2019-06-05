package stack
 
import (
	. "github.com/zrcoder/dsGo"
	base "github.com/zrcoder/dsGo/base/stack"
	"sync"
)
 
type Stack struct {
	lock  sync.RWMutex
	stack *base.Stack
}
 
func New() *Stack {
	return &Stack{stack: base.New()}
}
 
func (s *Stack) Push(item Any) {
	s.lock.Lock()
	s.stack.Push(item)
	s.lock.Unlock()
}
 
// Deletes top of a stack and returns it
func (s *Stack) Pop() Any {
	s.lock.Lock()
	item := s.stack.Pop()
	s.lock.Unlock()
	return item
}
 
// returns top of a stack without deletion
func (s *Stack) Peek() Any {
	s.lock.RLock()
	item := s.stack.Peek()
	s.lock.RUnlock()
	return item
}
