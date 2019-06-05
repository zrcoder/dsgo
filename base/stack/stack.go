package stack
 
import . "github.com/zrcoder/dsGo"
 
type stackItem struct {
	item Any
	next *stackItem
}
 
type Stack struct {
	peek  *stackItem
	depth uint64
}
 
func New() *Stack {
	return &Stack{}
}
 
func (s *Stack) Push(item Any) {
	s.peek = &stackItem{item: item, next: s.peek}
	s.depth ++
}
 
// Deletes top of a stack and returns it
func (s *Stack) Pop() Any {
	if s.depth == 0 {
		return nil
	}
	item := s.peek.item
	s.peek = s.peek.next
	s.depth --
	return item
}
 
// returns top of a stack without deletion
func (s *Stack) Peek() Any {
	if s.depth == 0 {
		return nil
	}
	return s.peek.item
}
