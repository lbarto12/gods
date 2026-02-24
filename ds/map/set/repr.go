package set

import "fmt"

func (set *Set[T]) String() string {
	if set.Len() == 0 {
		return "{}"
	}
	res := "{ "
	for key := range set.items {
		res += fmt.Sprintf("%v, ", key)
	}
	return res + "\b\b }"
}
