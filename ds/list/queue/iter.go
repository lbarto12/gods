package queue

func (q *Queue[T]) Iter() func(yield func(T) bool) {
	return q.items.Iter()
}

func (q *Queue[T]) IterPop() func(yield func(T) bool) {
	return q.items.IterPop()
}
