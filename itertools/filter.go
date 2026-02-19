package itertools

func Filter[T any](seq []T, lambda func(e T) bool) []T {
	res := []T{}

	for _, val := range seq {
		if lambda(val) {
			res = append(res, val)
		}
	}

	return res
}
