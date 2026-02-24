package bst

import "github.com/lbarto12/gods/common/godserr"

// Internal
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

// External
func (bst *BST[T]) Search(key T) (*T, error) {
	node := bst.find_ptr(&bst.root, key)
	if *node == nil {
		return nil, godserr.ErrKeyNotFound
	}
	return &(*node).val, nil
}
