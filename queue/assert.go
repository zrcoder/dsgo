package queue

import (
	"github.com/zrcoder/dsgo"
	"github.com/zrcoder/dsgo/list"
)

var _ dsgo.Container[int] = (*Queue[int])(nil)

func (q *Queue[T]) Len() int { return q.list.Len() }

func (q *Queue[T]) Empty() bool { return q.list.Empty() }

func (q *Queue[T]) Values() []T { return q.list.Values() }

func (q *Queue[T]) Clear() { q.list = list.New[T]() }
