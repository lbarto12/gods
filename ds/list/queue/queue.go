package queue

import "github.com/lbarto12/gods/ds/list/linkedlist"

type Queue[T any] struct {
	items *linkedlist.List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: linkedlist.NewList[T](),
	}
}

func (q *Queue[T]) String() string {
	return q.items.String()
}

func (q *Queue[T]) Push(val T) T {
	return q.items.PushFront(val)
}

func (q *Queue[T]) Pop() (*T, error) {
	return q.items.PopBack()
}

// Inherits

// control
func (q *Queue[T]) Copy() *Queue[T] {
	return &Queue[T]{
		items: q.items.Copy(),
	}
}

// data
func (q *Queue[T]) DumpSlice() []T {
	return q.items.DumpSlice()
}

func (q *Queue[T]) Len() uint64 {
	return q.items.Len()
}

// iterator
func (q *Queue[T]) Iter() func(yield func(T) bool) {
	return q.items.Iter()
}

func (q *Queue[T]) IterPop() func(yield func(T) bool) {
	return q.items.IterPop()
}
