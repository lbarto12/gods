package stack

func (stk *Stack[T]) String() string {
	return stk.items.String()
}
