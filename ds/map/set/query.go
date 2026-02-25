package set

func (set *Set[T]) IsDisjoint(other *Set[T]) bool {
	return set.Insersection(other).Len() == 0
}

func (set *Set[T]) IsSubset(other *Set[T]) bool {
	if set.Len() > other.Len() {
		return false
	}

	for k := range set.items {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}

func (set *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(set)
}
