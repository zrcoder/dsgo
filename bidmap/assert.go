package bidmap

import "github.com/zrcoder/dsgo"

var _ dsgo.BidMap[int, string] = (*Map[int, string])(nil)
