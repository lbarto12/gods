package set

func (set *Set[T]) Len() uint64 {
	return uint64(len(set.items))
}
