package set

func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.items[item] = struct{}{}
	}
}

func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(set.items, item)
	}
}
