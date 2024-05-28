package redblacktree

import (
	"cmp"

	"github.com/zrcoder/dsgo"
)

// Tree holds elements of the red-black tree
type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	size       int
	Comparator dsgo.Comparator[K]
}

// New instantiates a red-black tree with the built-in comparator for K
func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return NewWith[K, V](cmp.Compare[K])
}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith[K comparable, V any](comparator dsgo.Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{Comparator: comparator}
}

// Put inserts node into the tree or update the node's value if the key exsited.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Put(key K, value V) {
	insert := func(node, parent *Node[K, V]) {
		node.Parent = parent
		t.insertCase1(node)
		t.size++
	}
	if t.Root == nil {
		t.Comparator(key, key) // assert key is of comparator's type for initial tree
		t.Root = newNode(key, value)
		insert(t.Root, nil)
		return
	}
	cur := t.Root
	for {
		cmp := t.Comparator(key, cur.Key)
		switch {
		case cmp == 0:
			cur.Value = value
			return
		case cmp < 0:
			if cur.Left == nil {
				cur.Left = newNode(key, value)
				insert(cur.Left, cur)
				return
			}
			cur = cur.Left
		case cmp > 0:
			if cur.Right == nil {
				cur.Right = newNode(key, value)
				insert(cur.Right, cur)
				return
			}
			cur = cur.Right
		}
	}
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Get(key K) (value V, found bool) {
	node := t.GetNode(key)
	if node == nil {
		return
	}
	return node.Value, true
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) GetNode(key K) *Node[K, V] {
	cur := t.Root
	for cur != nil {
		cmp := t.Comparator(key, cur.Key)
		switch {
		case cmp == 0:
			return cur
		case cmp < 0:
			cur = cur.Left
		case cmp > 0:
			cur = cur.Right
		}
	}
	return nil
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Remove(key K) {
	node := t.GetNode(key)
	if node == nil {
		return
	}

	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	var child *Node[K, V]
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			t.deleteCase1(node)
		}
		t.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	t.size--
}

// Keys returns all keys in-order
func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, 0, t.size)
	t.Inorder(func(key K, value V) {
		keys = append(keys, key)
	})
	return keys
}

// Inorer travels the tree in-order with a handler.
func (t *Tree[K, V]) Inorder(handler func(key K, value V)) {
	t.Root.Inorder(handler)
}

// Left returns the left-most (min) node or nil if tree is empty.
func (t *Tree[K, V]) Left() *Node[K, V] {
	if t.Root == nil {
		return nil
	}
	cur := t.Root
	for cur != nil && cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// Right returns the right-most (max) node or nil if tree is empty.
func (t *Tree[K, V]) Right() *Node[K, V] {
	if t.Root == nil {
		return nil
	}
	cur := t.Root
	for cur != nil && cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool) {
	for cur := t.Root; cur != nil; {
		cmp := t.Comparator(key, cur.Key)
		switch {
		case cmp == 0:
			return cur, true
		case cmp < 0:
			cur = cur.Left
		case cmp > 0:
			floor, found = cur, true
			cur = cur.Right
		}
	}
	return
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[K, V]) Ceiling(key K) (ceiling *Node[K, V], found bool) {
	for cur := t.Root; cur != nil; {
		cmp := t.Comparator(key, cur.Key)
		switch {
		case cmp == 0:
			return cur, true
		case cmp < 0:
			ceiling, found = cur, true
			cur = cur.Left
		case cmp > 0:
			cur = cur.Right
		}
	}
	return
}

// String returns a string representation of container
func (t *Tree[K, V]) String() string {
	str := "RedBlackTree\n"
	if !t.Empty() {
		output(t.Root, "", true, &str)
	}
	return str
}

func (t *Tree[K, V]) rotateLeft(node *Node[K, V]) {
	right := node.Right
	t.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (t *Tree[K, V]) rotateRight(node *Node[K, V]) {
	left := node.Left
	t.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (t *Tree[K, V]) replaceNode(old *Node[K, V], new *Node[K, V]) {
	if old.Parent == nil {
		t.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (t *Tree[K, V]) insertCase1(node *Node[K, V]) {
	if node.Parent == nil {
		node.color = black
	} else {
		t.insertCase2(node)
	}
}

func (t *Tree[K, V]) insertCase2(node *Node[K, V]) {
	if nodeColor(node.Parent) == black {
		return
	}
	t.insertCase3(node)
}

func (t *Tree[K, V]) insertCase3(node *Node[K, V]) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		t.insertCase1(node.grandparent())
	} else {
		t.insertCase4(node)
	}
}

func (t *Tree[K, V]) insertCase4(node *Node[K, V]) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		t.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		t.rotateRight(node.Parent)
		node = node.Right
	}
	t.insertCase5(node)
}

func (t *Tree[K, V]) insertCase5(node *Node[K, V]) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		t.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		t.rotateLeft(grandparent)
	}
}

func (t *Tree[K, V]) deleteCase1(node *Node[K, V]) {
	if node.Parent == nil {
		return
	}
	t.deleteCase2(node)
}

func (t *Tree[K, V]) deleteCase2(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			t.rotateLeft(node.Parent)
		} else {
			t.rotateRight(node.Parent)
		}
	}
	t.deleteCase3(node)
}

func (t *Tree[K, V]) deleteCase3(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		t.deleteCase1(node.Parent)
	} else {
		t.deleteCase4(node)
	}
}

func (t *Tree[K, V]) deleteCase4(node *Node[K, V]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		t.deleteCase5(node)
	}
}

func (t *Tree[K, V]) deleteCase5(node *Node[K, V]) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		t.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		t.rotateLeft(sibling)
	}
	t.deleteCase6(node)
}

func (t *Tree[K, V]) deleteCase6(node *Node[K, V]) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		t.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		t.rotateRight(node.Parent)
	}
}
