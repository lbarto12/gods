package set

import (
	"slices"
)

func (set *Set[T]) Contains(items ...T) bool {
	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}
	return true
}

func (set *Set[T]) Union(other *Set[T]) *Set[T] {
	s := set.Copy()
	s.Add(slices.Collect(other.Iter())...)
	return s
}

func (set *Set[T]) Merge(other *Set[T]) *Set[T] {
	set.Add(slices.Collect(other.Iter())...)
	return set
}

func (set *Set[T]) Insersection(other *Set[T]) *Set[T] {
	s := New[T]()
	for k := range set.Iter() {
		if other.Contains(k) {
			s.Add(k)
		}
	}

	return s
}

func (set *Set[T]) SymDiff(other *Set[T]) *Set[T] {
	s := set.Copy().Merge(other)

	for k := range set.Iter() {
		if set.Contains(k) && other.Contains(k) {
			s.Remove(k)
		}
	}

	return s
}
