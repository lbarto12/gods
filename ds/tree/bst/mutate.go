package bst

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
