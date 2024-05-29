package bidmap

import (
	"fmt"
	"strings"
)

type Map[K, V comparable] struct {
	kv map[K]V
	vk map[V]K
}

func New[K, V comparable]() *Map[K, V] {
	return &Map[K, V]{kv: map[K]V{}, vk: map[V]K{}}
}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.kv[key]
	return
}

func (m *Map[K, V]) GetKey(value V) (key K, ok bool) {
	key, ok = m.vk[value]
	return
}

func (m *Map[K, V]) Put(key K, value V) {
	if v, ok := m.kv[key]; ok {
		delete(m.vk, v)
	}
	if k, ok := m.vk[value]; ok {
		delete(m.kv, k)
	}
	m.kv[key] = value
	m.vk[value] = key
}

func (m *Map[K, V]) Remove(key K) {
	value, ok := m.kv[key]
	if !ok {
		return
	}
	delete(m.kv, key)
	delete(m.vk, value)
}

func (m *Map[K, V]) RemoveValue(value V) {
	key, ok := m.vk[value]
	if !ok {
		return
	}
	delete(m.vk, value)
	delete(m.kv, key)
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "BidMap\nmap["
	kvs := make([]string, 0, len(m.kv))
	for key, value := range m.kv {
		kvs = append(kvs, fmt.Sprintf("%v:%v", key, value))
	}
	return str + strings.Join(kvs, " ") + " ]"
}

func (m *Map[K, V]) Len() int    { return len(m.kv) }
func (m *Map[K, V]) Empty() bool { return len(m.kv) == 0 }

func (m *Map[K, V]) Keys() []K {
	res := make([]K, len(m.kv))
	for key := range m.kv {
		res = append(res, key)
	}
	return res
}

func (m *Map[K, V]) Values() []V {
	res := make([]V, len(m.kv))
	for _, value := range m.kv {
		res = append(res, value)
	}
	return res
}

func (m *Map[K, V]) Clear() {
	clear(m.kv)
	clear(m.vk)
}
