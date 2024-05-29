package lfucache

import "testing"

func Test(t *testing.T) {
	opers := [][]any{
		{"new", 3},
		{"put", 1, 1},
		{"put", 2, 2},
		{"put", 3, 2},
		{"put", 2, 4},
		{"put", 3, 5},
		{"get", 2, 4, true},
		{"put", 4, 4},
		{"get", 1, 0, false},
	}
	test(t, opers)
}

func test(t *testing.T, opers [][]any) {
	var cache *Cache[int, int]
	for i, oper := range opers {
		switch oper[0] {
		case "new":
			cache = New[int, int](oper[1].(int))
		case "put":
			cache.Put(oper[1].(int), oper[2].(int))
		case "get":
			got, ok := cache.Get(oper[1].(int))
			if got != oper[2].(int) || ok != oper[3].(bool) {
				t.Errorf("operation %d: want %d %v, got %d %v",
					i, oper[2].(int), oper[3].(bool), got, ok)
			}
		}
	}
}
