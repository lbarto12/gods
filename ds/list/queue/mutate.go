package queue

func (q *Queue[T]) Push(val T) T {
	return q.items.PushFront(val)
}

func (q *Queue[T]) Pop() (*T, error) {
	return q.items.PopBack()
}
