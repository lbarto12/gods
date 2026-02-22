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

// Inherits

// control
func (stk *Stack[T]) Copy() *Stack[T] {
	return &Stack[T]{
		items: stk.items.Copy(),
	}
}

// data
func (stk *Stack[T]) DumpSlice() []T {
	return stk.items.DumpSlice()
}

func (stk *Stack[T]) Len() uint64 {
	return stk.items.Len()
}

// iterator
func (stk *Stack[T]) Iter() func(yield func(T) bool) {
	return stk.items.Iter()
}

func (stk *Stack[T]) IterPop() func(yield func(T) bool) {
	return stk.items.IterPop()
}
