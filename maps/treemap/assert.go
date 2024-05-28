package treemap

import "github.com/zrcoder/dsgo/maps"

var _ maps.Map[string, int] = (*Map[string, int])(nil)
