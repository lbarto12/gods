package set

import (
	"maps"
	"slices"
)

func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}
	return true
}

func (set *Set[T]) Union(other Set[T]) *Set[T] {
	items_clone := maps.Clone(set.items)
	new_set := Set[T]{
		items: items_clone,
	}

	new_set.Add(slices.Collect(maps.Keys(other.items))...)
	return &new_set
}
