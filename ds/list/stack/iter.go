package stack

import "iter"

func (stk *Stack[T]) Iter() iter.Seq[T] {
	return stk.items.Iter()
}

func (stk *Stack[T]) IterPop() iter.Seq[T] {
	return stk.items.IterPop()
}
