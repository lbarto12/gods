package counter

import "math"

func (c Counter[T]) Get() map[T]uint64 {
	return c.counts
}

func (c Counter[T]) Max() T {
	var bst uint64 = 0
	var key T

	for k, v := range c.counts {
		if v > bst {
			key = k
			bst = v
		}
	}

	return key
}

func (c Counter[T]) Min() T {
	var bst uint64 = uint64(math.MaxUint64)
	var key T

	for k, v := range c.counts {
		if v < bst {
			key = k
			bst = v
		}
	}

	return key
}
