package stack

import "github.com/lbarto12/gods/ds/list/linkedlist"

type Stack[T any] struct {
	items *linkedlist.List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: linkedlist.New[T](),
	}
}
