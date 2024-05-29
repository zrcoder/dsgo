package lrucache

import (
	"github.com/zrcoder/dsgo/list"
	"github.com/zrcoder/dsgo/maps"
)

type Cache[K comparable, V any] struct {
	size int
	m    map[K]*list.Element[maps.Pair[K, V]]
	list *list.List[maps.Pair[K, V]]
}

func New[K comparable, V any](size int) *Cache[K, V] {
	if size < 1 {
		panic("cache size must more than 0")
	}
	return &Cache[K, V]{
		size: size,
		m:    make(map[K]*list.Element[maps.Pair[K, V]], size),
		list: list.New[maps.Pair[K, V]](),
	}
}

func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	if e, ok := c.m[key]; ok {
		c.list.MoveToFront(e)
		return e.Value.Value, true
	}
	return
}

func (c *Cache[K, V]) Put(key K, value V) {
	if e, ok := c.m[key]; ok {
		e.Value.Value = value
		c.list.MoveToFront(e)
		return
	}
	if c.list.Len() == c.size {
		pair := c.list.Remove(c.list.Back())
		delete(c.m, pair.Key)
	}
	e := c.list.PushFront(maps.Pair[K, V]{Key: key, Value: value})
	c.m[key] = e
}

func (c *Cache[K, V]) Len() int { return c.list.Len() }

func (c *Cache[K, V]) Empty() bool { return c.list.Empty() }

func (c *Cache[K, V]) Keys() []K {
	res := make([]K, 0, c.list.Len())
	for e := c.list.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.Key)
	}
	return res
}

func (c *Cache[K, V]) Values() []V {
	res := make([]V, 0, c.list.Len())
	for e := c.list.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.Value)
	}
	return res
}

func (c *Cache[K, V]) Clear() {
	clear(c.m)
	c.list.Clear()
}
