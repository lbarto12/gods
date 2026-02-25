package queue

func (q *Queue[T]) Peek() (*T, error) {
	return q.items.Tail()
}
