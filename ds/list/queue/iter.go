package queue

import "iter"

func (q *Queue[T]) Iter() iter.Seq[T] {
	return q.items.Iter()
}

func (q *Queue[T]) IterPop() iter.Seq[T] {
	return q.items.IterPop()
}
