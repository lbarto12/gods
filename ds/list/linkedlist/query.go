package linkedlist

import "github.com/lbarto12/gods/common/godserr"

func (lst *List[T]) Tail() (*T, error) {
	if lst.Len() == 0 {
		return nil, godserr.ErrNoItems
	}
	val := lst.tail.val
	return &val, nil
}

func (lst *List[T]) Head() (*T, error) {
	if lst.Len() == 0 {
		return nil, godserr.ErrNoItems
	}
	val := lst.head.val
	return &val, nil
}
