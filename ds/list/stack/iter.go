package stack

func (stk *Stack[T]) Iter() func(yield func(T) bool) {
	return stk.items.Iter()
}

func (stk *Stack[T]) IterPop() func(yield func(T) bool) {
	return stk.items.IterPop()
}
