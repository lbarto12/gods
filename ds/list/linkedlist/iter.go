package linkedlist

func (lst *List[T]) Iter() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		cpy := lst.Copy()
		for v := range cpy.IterPop() {
			if !yield(v) {
				return
			}
		}
	}
}

func (lst *List[T]) IterPop() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for true {
			elem, err := lst.PopFront()
			if err != nil {
				break
			}
			if !yield(*elem) {
				return
			}
		}
	}
}
