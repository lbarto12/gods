package linkedlist

import (
	"github.com/lbarto12/gods/common/godserr"
)

func (lst *List[T]) PushFront(val T) T {
	lst.head = &list_node[T]{
		val:  val,
		next: lst.head,
		prev: nil,
	}
	if lst.head.next != nil {
		lst.head.next.prev = lst.head
	} else {
		lst.tail = lst.head
	}
	lst.size++
	return val
}

func (lst *List[T]) PopFront() (*T, error) {
	if lst.head == nil {
		return nil, godserr.ErrNoItems
	}

	store := lst.head
	lst.head = lst.head.next
	if lst.head != nil {
		lst.head.prev = nil
	} else {
		lst.tail = nil
	}
	lst.size--
	return &store.val, nil
}

func (lst *List[T]) PushBack(val T) T {
	if lst.head == nil {
		lst.head = &list_node[T]{
			val:  val,
			next: nil,
			prev: nil,
		}
		lst.tail = lst.head
		lst.size++
		return val
	}

	lst.tail.next = &list_node[T]{
		val:  val,
		next: nil,
		prev: lst.tail,
	}
	lst.tail = lst.tail.next
	lst.size++
	return val
}

func (lst *List[T]) PopBack() (*T, error) {
	if lst.tail == nil {
		return nil, godserr.ErrNoItems
	}

	store := lst.tail.val
	lst.tail = lst.tail.prev
	if lst.tail != nil {
		lst.tail.next = nil
	} else {
		lst.head = nil
	}
	lst.size--
	return &store, nil
}
