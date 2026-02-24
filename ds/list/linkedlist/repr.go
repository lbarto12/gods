package linkedlist

import "fmt"

func (lst *List[T]) String() string {
	if lst.head == nil {
		return "[]"
	}
	res := "[ "
	ref := lst.head
	for ref != nil {
		res += fmt.Sprintf("%v, ", ref.val)
		ref = ref.next
	}
	return res + "\b\b ]"
}
