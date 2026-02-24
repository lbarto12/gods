package stack

func (stk *Stack[T]) Copy() *Stack[T] {
	return &Stack[T]{
		items: stk.items.Copy(),
	}
}
