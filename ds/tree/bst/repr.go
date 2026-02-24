package bst

import "fmt"

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
