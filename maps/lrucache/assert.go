package lrucache

import "github.com/zrcoder/dsgo/maps"

var _ maps.Cache[int, string] = (*Cache[int, string])(nil)
