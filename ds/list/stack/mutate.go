package stack

func (stk *Stack[T]) Push(val T) T {
	return stk.items.PushFront(val)
}

func (stk *Stack[T]) Pop() (*T, error) {
	return stk.items.PopFront()
}
