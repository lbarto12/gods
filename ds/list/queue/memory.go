package queue

func (q *Queue[T]) Copy() *Queue[T] {
	return &Queue[T]{
		items: q.items.Copy(),
	}
}
