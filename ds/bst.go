package ds

import (
	"cmp"
	"fmt"

	"github.com/lbarto12/gods/util"
)

type bst_node[T any] struct {
	val   T
	left  *bst_node[T]
	right *bst_node[T]
}

type BST[T any] struct {
	root       *bst_node[T]
	comparator util.Comparator[T]
}

func NewBST[T cmp.Ordered]() *BST[T] {
	return NewBSTWithComparator(cmp.Compare[T])
}

func NewBSTWithComparator[T cmp.Ordered](comparator util.Comparator[T]) *BST[T] {
	return &BST[T]{
		root:       nil,
		comparator: comparator,
	}
}

func (bst *BST[T]) insert_driver(node **bst_node[T], val T) T {
	if *node == nil {
		*node = &bst_node[T]{
			val: val,
		}
		return val
	}

	comp := bst.comparator(val, (*node).val)
	switch {
	case comp < 0:
		return bst.insert_driver(&(*node).left, val)
	case comp > 0:
		return bst.insert_driver(&(*node).right, val)
	case comp == 0: // change existing
		(*node).val = val
		return (*node).val
	}

	return val
}

// Methods
func (bst *BST[T]) Insert(val T) T {
	return bst.insert_driver(&bst.root, val)
}

// data

// string
func (bst *BST[T]) to_string(node *bst_node[T]) string {
	if node == nil {
		return ""
	}
	return bst.to_string(node.left) + fmt.Sprintf("%v", node.val) + ", " + bst.to_string(node.right)
}

func (bst *BST[T]) String() string {
	repr := bst.to_string(bst.root)
	if repr == "" {
		return "[]"
	}

	return fmt.Sprintf("[ %s\b\b ]", repr)
}
