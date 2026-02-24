package counter

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
