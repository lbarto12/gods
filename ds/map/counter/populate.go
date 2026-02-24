package counter

func (c Counter[T]) count(seq []T) map[T]uint64 {
	res := make(map[T]uint64)

	for _, elem := range seq {
		res[elem]++
	}

	return res
}
