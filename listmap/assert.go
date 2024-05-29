package listmap

import "github.com/zrcoder/dsgo"

var _ dsgo.Map[string, int] = (*Map[string, int])(nil)
