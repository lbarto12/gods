package set

import "slices"

func (set *Set[T]) Copy() *Set[T] {
	return New(slices.Collect(set.Iter())...)
}
