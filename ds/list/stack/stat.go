package stack

func (stk *Stack[T]) Len() uint64 {
	return stk.items.Len()
}
