package queue

func (q *Queue[T]) String() string {
	return q.items.String()
}
