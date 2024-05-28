package redblacktree

import "github.com/zrcoder/dsgo"

var _ dsgo.Container[int] = (*Tree[string, int])(nil)

// Empty returns true if tree does not contain any nodes
func (tree *Tree[K, V]) Empty() bool { return tree.size == 0 }

// Len returns number of nodes in the tree.
func (tree *Tree[K, V]) Len() int { return tree.size }

// Values returns all values in-order based on the key.
func (t *Tree[K, V]) Values() []V {
	values := make([]V, 0, t.size)
	t.Inorder(func(key K, value V) {
		values = append(values, value)
	})
	return values
}

// Clear removes all nodes from the tree.
func (tree *Tree[K, V]) Clear() {
	tree.Root = nil
	tree.size = 0
}
