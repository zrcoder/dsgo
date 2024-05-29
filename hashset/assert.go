package hashset

import (
	"github.com/zrcoder/dsgo"
)

// Assert Set implementation
var _ dsgo.Set[int] = (*Set[int])(nil)
