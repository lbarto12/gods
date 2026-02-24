package counter

import (
	"math"
)

type Counter[T comparable] struct {
	counts map[T]uint64
}

func NewCounter[T comparable](seq []T) *Counter[T] {
	c := Counter[T]{
		counts: make(map[T]uint64),
	}
	c.counts = c.count(seq)
	return &c
}

func (c Counter[T]) count(seq []T) map[T]uint64 {
	res := make(map[T]uint64)

	for _, elem := range seq {
		res[elem]++
	}

	return res
}

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
