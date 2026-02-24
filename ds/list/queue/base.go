package queue

import "github.com/lbarto12/gods/ds/list/linkedlist"

type Queue[T any] struct {
	items *linkedlist.List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: linkedlist.New[T](),
	}
}
