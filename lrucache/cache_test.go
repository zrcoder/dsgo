package lrucache

import "testing"

func Test(t *testing.T) {
	opers := [][]any{
		{"new", 2},
		{"put", 1, 1},
		{"put", 2, 2},
		{"get", 1, 1, true},
		{"put", 3, 3},
		{"get", 2, 0, false},
		{"put", 4, 4},
		{"get", 1, 0, false},
		{"get", 3, 3, true},
		{"get", 4, 4, true},
	}
	test(t, opers)
}

func test(t *testing.T, opers [][]any) {
	t.Helper()
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
