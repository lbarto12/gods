package ds

type Stack[T any] struct {
	items *List[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: NewList[T](),
	}
}

func (stk *Stack[T]) String() string {
	return stk.items.String()
}

func (stk *Stack[T]) Push(val T) T {
	return stk.items.PushFront(val)
}

func (stk *Stack[T]) Pop() (*T, error) {
	return stk.items.PopFront()
}
