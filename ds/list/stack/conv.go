package stack

func (stk *Stack[T]) DumpSlice() []T {
	return stk.items.DumpSlice()
}
