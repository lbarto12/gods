package itertools

func Zip[T any](args ...[]T) [][]T {
	res := [][]T{}

	for i, v := range args[0] {
		res = append(res, []T{v})
		for _, arr := range args[1:] {
			res[i] = append(res[i], arr[i])
		}
	}

	return res
}
