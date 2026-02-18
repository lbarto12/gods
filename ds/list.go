package ds

import (
	"errors"
	"fmt"
)

type list_node[T any] struct {
	val  T
	next *list_node[T]
	prev *list_node[T]
}

type List[T any] struct {
	head *list_node[T]
	tail *list_node[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

var ( // Errors
	ErrNoItems = errors.New("no items found")
)

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
	return val
}

func (lst *List[T]) PopFront() (*T, error) {
	if lst.head == nil {
		return nil, ErrNoItems
	}

	store := lst.head
	lst.head = lst.head.next
	if lst.head != nil {
		lst.head.prev = nil
	} else {
		lst.tail = nil
	}
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
		return val
	}

	lst.tail.next = &list_node[T]{
		val:  val,
		next: nil,
		prev: lst.tail,
	}
	lst.tail = lst.tail.next
	return val
}

func (lst *List[T]) PopBack() (*T, error) {
	if lst.tail == nil {
		return nil, ErrNoItems
	}

	store := lst.tail.val
	lst.tail = lst.tail.prev
	if lst.tail != nil {
		lst.tail.next = nil
	} else {
		lst.head = nil
	}

	return &store, nil
}
