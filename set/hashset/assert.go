package hashset

import (
	"github.com/zrcoder/dsgo/set"
)

// Assert Set implementation
var _ set.Set[int] = (*Set[int])(nil)
