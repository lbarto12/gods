package queue

func (q *Queue[T]) Len() uint64 {
	return q.items.Len()
}
