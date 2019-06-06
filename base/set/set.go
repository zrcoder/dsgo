package set

import . "github.com/zrcoder/dsGo"

type Set map[Any]Empty

func New() Set {
	return make(map[Any]Empty)
}

func NewWithCapacity(c int) Set  {
	return make(map[Any]Empty, c)
}

func (s Set) Add(item Any) {
	s[item] = Empty{}
}

func (s Set) Delete(item Any) {
	delete(s, item)
}

func (s Set) Has(item Any) bool {
	_, ok := s[item]
	return ok
}

func (s Set) Count() int {
	return len(s)
}

func (s Set) AllItems() []Any {
	r := make([]Any, s.Count())
	for k := range s {
		r = append(r, k)
	}
	return r
}
