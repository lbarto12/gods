package linkedlist

func (lst *List[T]) DumpSlice() []T {
	temp := lst.Copy()
	res := []T{}
	for val := range temp.IterPop() {
		res = append(res, val)
	}
	return res
}
