package ds

import (
	"strings"
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

// external
func StringArray(s string) []string {
	return strings.Split(s, "")
}
