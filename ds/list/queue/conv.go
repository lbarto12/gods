package queue

func (q *Queue[T]) DumpSlice() []T {
	return q.items.DumpSlice()
}
