package redblacktree

import "fmt"

type color bool

const (
	black, red color = true, false
)

// Node is a single element within the tree
type Node[K comparable, V any] struct {
	Key    K
	Value  V
	color  color
	Left   *Node[K, V]
	Right  *Node[K, V]
	Parent *Node[K, V]
}

func newNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{Key: key, Value: value, color: red}
}

// Inorer travels the node as root in-order with a handler.
// Morris traversal used,
// the time complex is O(n), n is the total nodes in the tree,
// and the space complex is O(1)
func (n *Node[K, V]) Inorder(handler func(key K, value V)) {
	cur := n
	var pre *Node[K, V]
	for cur != nil {
		if cur.Left == nil {
			handler(cur.Key, cur.Value)
			cur = cur.Right
			continue
		}
		pre = cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			handler(cur.Key, cur.Value)
			cur = cur.Right
		}
	}
}

// Size returns the number of elements stored in the subtree.
// Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.
func (n *Node[K, V]) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.Left.Size() + n.Right.Size()
}

func (n *Node[K, V]) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func output[K comparable, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

func (n *Node[K, V]) grandparent() *Node[K, V] {
	if n != nil && n.Parent != nil {
		return n.Parent.Parent
	}
	return nil
}

func (n *Node[K, V]) uncle() *Node[K, V] {
	if n == nil || n.Parent == nil || n.Parent.Parent == nil {
		return nil
	}
	return n.Parent.sibling()
}

func (n *Node[K, V]) sibling() *Node[K, V] {
	if n == nil || n.Parent == nil {
		return nil
	}
	if n == n.Parent.Left {
		return n.Parent.Right
	}
	return n.Parent.Left
}

func (n *Node[K, V]) maximumNode() *Node[K, V] {
	if n == nil {
		return nil
	}
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func nodeColor[K comparable, V any](node *Node[K, V]) color {
	if node == nil {
		return black
	}
	return node.color
}
