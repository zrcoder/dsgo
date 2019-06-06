package set
 
import (
	"sync"
 
	. "github.com/zrcoder/dsGo"
	base "github.com/zrcoder/dsGo/base/set"
)
 
type Set struct {
	lock sync.RWMutex
	set  base.Set
}
 
func New() *Set {
	s := &Set{}
	s.set = base.New()
	return s
}
 
func NewWithCapacity(c int) *Set {
	s := &Set{}
	s.set = base.NewWithCapacity(c)
	return s
}
 
func (s *Set) Add(item Any) {
	s.lock.Lock()
	s.set.Add(item)
	s.lock.Unlock()
}
 
func (s *Set) Delete(item Any) {
	s.lock.Lock()
	s.set.Delete(item)
	s.lock.Unlock()
}
 
func (s *Set) Has(item Any) bool {
	s.lock.RLock()
	exist := s.set.Has(item)
	s.lock.RUnlock()
	return exist
}
 
func (s *Set) Count() int {
	s.lock.RLock()
	length := s.set.Count()
	s.lock.RUnlock()
	return length
}
 
func (s *Set) AllItems() []Any {
	s.lock.RLock()
	items := s.set.AllItems()
	s.lock.RUnlock()
	return items
}
