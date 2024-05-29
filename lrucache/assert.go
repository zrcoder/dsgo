package lrucache

import "github.com/zrcoder/dsgo"

var _ dsgo.Cache[int, string] = (*Cache[int, string])(nil)
