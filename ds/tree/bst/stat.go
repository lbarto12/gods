package bst

import "github.com/lbarto12/gods/common/godserr"

func (bst *BST[T]) Max() (*T, error) {
	if bst.root == nil {
		return nil, godserr.ErrNoItems
	}
	node := bst.find_max(bst.root)
	return &node.val, nil
}

func (bst *BST[T]) Min() (*T, error) {
	if bst.root == nil {
		return nil, godserr.ErrNoItems
	}
	node := bst.find_min(bst.root)
	return &node.val, nil
}
