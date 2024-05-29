package lfucache

import (
	"github.com/zrcoder/dsgo/list"
)

type Item[K comparable, V any] struct {
	Key   K
	Value V
	Freq  int
}

type Cache[K comparable, V any] struct {
	keyElements map[K]*list.Element[Item[K, V]]
	freqLists   map[int]*list.List[Item[K, V]]
	size        int
	minFreq     int
}

func New[K comparable, V any](size int) *Cache[K, V] {
	if size < 1 {
		panic("the cache size must more than 0")
	}
	return &Cache[K, V]{
		keyElements: make(map[K]*list.Element[Item[K, V]], size),
		freqLists:   make(map[int]*list.List[Item[K, V]]),
		size:        size,
	}
}

func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	if element, ok := c.keyElements[key]; ok {
		c.increseFreq(element)
		return element.Value.Value, true
	}
	return
}

func (c *Cache[K, V]) Put(key K, value V) {
	if element, ok := c.keyElements[key]; ok {
		element.Value.Value = value
		c.increseFreq(element)
		return
	}

	if len(c.keyElements) == c.size {
		// remove the min freq element
		list := c.freqLists[c.minFreq]
		key := list.Remove(list.Back()).Key
		delete(c.keyElements, key)
		// will update c.minFreq later
	}

	if c.freqLists[1] == nil {
		c.freqLists[1] = list.New[Item[K, V]]()
	}
	item := Item[K, V]{Freq: 1, Value: value, Key: key}
	c.keyElements[key] = c.freqLists[1].PushFront(item)
	c.minFreq = 1
}

func (c *Cache[K, V]) increseFreq(element *list.Element[Item[K, V]]) {
	item := element.Value
	oldList := c.freqLists[item.Freq]
	oldList.Remove(element)
	if oldList.Len() == 0 && c.minFreq == item.Freq {
		c.minFreq++
	}
	item.Freq++
	if _, ok := c.freqLists[item.Freq]; !ok {
		c.freqLists[item.Freq] = list.New[Item[K, V]]()
	}
	newList := c.freqLists[item.Freq]
	element = newList.PushFront(item)
	c.keyElements[item.Key] = element
}

func (c *Cache[K, V]) Len() int { return len(c.keyElements) }

func (c *Cache[K, V]) Empty() bool { return len(c.keyElements) == 0 }

func (c *Cache[K, V]) Keys() []K {
	res := make([]K, 0, len(c.keyElements))
	for key := range c.keyElements {
		res = append(res, key)
	}
	return res
}

func (c *Cache[K, V]) Values() []V {
	res := make([]V, 0, len(c.keyElements))
	for _, element := range c.keyElements {
		res = append(res, element.Value.Value)
	}
	return res
}

func (c *Cache[K, V]) Clear() {
	clear(c.keyElements)
	clear(c.freqLists)
	c.minFreq = 0
}
