package bst

import (
	"cmp"
	"errors"
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

// errors
var (
	ErrKeyNotFound = errors.New("key not found in tree")
	ErrTreeEmpty   = errors.New("tree is empty")
)

// Methods

func (bst *BST[T]) find_ptr(node **bst_node[T], key T) **bst_node[T] {
	if *node == nil {
		return node
	}

	comp := bst.comparator(key, (*node).val)
	switch {
	case comp < 0:
		return bst.find_ptr(&(*node).left, key)
	case comp > 0:
		return bst.find_ptr(&(*node).right, key)
	}

	return node
}

func (bst *BST[T]) find_max(node *bst_node[T]) *bst_node[T] {
	if node.right == nil {
		return node
	}
	return bst.find_max(node.right)
}

func (bst *BST[T]) find_min(node *bst_node[T]) *bst_node[T] {
	if node.left == nil {
		return node
	}
	return bst.find_min(node.left)
}

// INSERT
func (bst *BST[T]) Insert(val T) T {
	node := bst.find_ptr(&bst.root, val)
	if *node == nil {
		*node = &bst_node[T]{
			val: val,
		}
	} else {
		(*node).val = val
	}
	return (*node).val
}

// REMOVE
func (bst *BST[T]) remove_driver(node **bst_node[T], val T) error {
	del := bst.find_ptr(node, val)
	if *del == nil {
		return ErrKeyNotFound
	}

	if (*del).left == nil { // No children, or one node on right
		(*del) = (*del).right // Fold right upwards
	} else if (*del).right == nil { // Only left child
		(*del) = (*del).left // Fold left upwards
	} else { // Two children
		mx_node := bst.find_max((*del).left) // will not be null
		(*del).val = mx_node.val
		return bst.remove_driver(&(*del).left, mx_node.val)
	}
	return nil
}

func (bst *BST[T]) Remove(val T) error {
	return bst.remove_driver(&bst.root, val)
}

// data
func (bst *BST[T]) Max() (*T, error) {
	if bst.root == nil {
		return nil, ErrTreeEmpty
	}
	node := bst.find_max(bst.root)
	return &node.val, nil
}

func (bst *BST[T]) Min() (*T, error) {
	if bst.root == nil {
		return nil, ErrTreeEmpty
	}
	node := bst.find_min(bst.root)
	return &node.val, nil
}

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
