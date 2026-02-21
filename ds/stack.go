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
func (q *Stack[T]) Copy() *Stack[T] {
	return &Stack[T]{
		items: q.items.Copy(),
	}
}

// data
func (q *Stack[T]) DumpSlice() []T {
	return q.items.DumpSlice()
}

// iterator
func (q *Stack[T]) Iter() func(yield func(T) bool) {
	return q.items.Iter()
}

func (q *Stack[T]) IterPop() func(yield func(T) bool) {
	return q.items.IterPop()
}
