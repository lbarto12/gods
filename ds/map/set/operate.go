package set

import (
	"maps"
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
	items_clone := maps.Clone(set.items)
	new_set := Set[T]{
		items: items_clone,
	}

	new_set.Add(slices.Collect(maps.Keys(other.items))...)
	return &new_set
}

func (set *Set[T]) Merge(other *Set[T]) *Set[T] {
	for k := range other.items {
		set.items[k] = struct{}{}
	}

	return set
}

func (set *Set[T]) Insersection(other *Set[T]) *Set[T] {
	new_data := map[T]struct{}{}
	for k := range set.items {
		if other.Contains(k) {
			new_data[k] = struct{}{}
		}
	}

	new_set := Set[T]{
		items: new_data,
	}

	return &new_set
}

func (set *Set[T]) SymDiff(other *Set[T]) *Set[T] {
	s := New(slices.Collect(set.Iter())...)
	s.Merge(other)

	for k := range set.items {
		if set.Contains(k) && other.Contains(k) {
			s.Remove(k)
		}
	}

	return s
}
