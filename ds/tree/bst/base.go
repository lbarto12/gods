package bst

import (
	"cmp"

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
