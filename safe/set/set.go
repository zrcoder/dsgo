/*
	A set can store unique values, without any particular order.
*/
package set
 
import (
	"sync"
 
	. "github.com/zrcoder/dsGo"
	base "github.com/zrcoder/dsGo/base/set"
)
 
type Set struct {
	lock sync.RWMutex
	inner  base.Set
}
 
func New() *Set {
	s := &Set{}
	s.inner = base.New()
	return s
}
 
func NewWithCapacity(c int) *Set {
	s := &Set{}
	s.inner = base.NewWithCapacity(c)
	return s
}
 
func (s *Set) Add(item Any) {
	s.lock.Lock()
	s.inner.Add(item)
	s.lock.Unlock()
}
 
func (s *Set) Delete(item Any) {
	s.lock.Lock()
	s.inner.Delete(item)
	s.lock.Unlock()
}
 
func (s *Set) Has(item Any) bool {
	s.lock.RLock()
	exist := s.inner.Has(item)
	s.lock.RUnlock()
	return exist
}
 
func (s *Set) Size() int {
	s.lock.RLock()
	length := s.inner.Size()
	s.lock.RUnlock()
	return length
}
 
func (s *Set) AllItems() []Any {
	s.lock.RLock()
	items := s.inner.AllItems()
	s.lock.RUnlock()
	return items
}
