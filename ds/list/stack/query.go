package stack

func (stk *Stack[T]) Peek() (*T, error) {
	return stk.items.Head()
}
