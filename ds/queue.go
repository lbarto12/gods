package ds

type Queue[T any] struct {
	items *List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: NewList[T](),
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

// tst
