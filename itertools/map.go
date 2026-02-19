package itertools

func Map[T any, K any](seq []T, lambda func(e T) K) []K {
	res := []K{}

	for _, val := range seq {
		res = append(res, lambda(val))
	}

	return res
}
