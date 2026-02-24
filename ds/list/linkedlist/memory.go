package linkedlist

func copy_driver[T any](o *list_node[T]) (*list_node[T], *list_node[T]) {
	if o == nil {
		return nil, nil
	}
	if o.next == nil {
		return o, o
	}
	nxt, tail := copy_driver(o.next)
	newnode := &list_node[T]{
		val:  o.val,
		next: nxt,
	}
	if newnode.next != nil {
		newnode.next.prev = newnode
	}

	return newnode, tail
}

func (lst *List[T]) Copy() *List[T] {
	head, tail := copy_driver(lst.head)
	return &List[T]{
		head: head,
		tail: tail,
		size: lst.size,
	}
}
