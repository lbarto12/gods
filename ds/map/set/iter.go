package set

import (
	"iter"
)

func (set *Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range set.items {
			if !yield(k) {
				return
			}
		}
	}
}
