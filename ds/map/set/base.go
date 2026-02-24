package set

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable](items ...T) *Set[T] {
	set := &Set[T]{
		items: map[T]struct{}{},
	}

	set.Add(items...)
	return set
}
