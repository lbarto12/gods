package linkedlist

type list_node[T any] struct {
	val  T
	next *list_node[T]
	prev *list_node[T]
}

type List[T any] struct {
	head *list_node[T]
	tail *list_node[T]
	size uint64
}

func New[T any]() *List[T] {
	return &List[T]{
		size: 0,
	}
}
