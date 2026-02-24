package linkedlist

import (
	"fmt"

	"github.com/lbarto12/gods/ds"
)

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

func NewList[T any]() *List[T] {
	return &List[T]{
		size: 0,
	}
}

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
	lst.size++
	return val
}

func (lst *List[T]) PopFront() (*T, error) {
	if lst.head == nil {
		return nil, ds.ErrNoItems
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
		return nil, ds.ErrNoItems
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

// mem control

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

// data
func (lst *List[T]) DumpSlice() []T {
	temp := lst.Copy()
	res := []T{}
	for val := range temp.IterPop() {
		res = append(res, val)
	}
	return res
}

func (lst *List[T]) Len() uint64 {
	return lst.size
}

// Iterators
func (lst *List[T]) Iter() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		cpy := lst.Copy()
		for v := range cpy.IterPop() {
			if !yield(v) {
				return
			}
		}
	}
}

func (lst *List[T]) IterPop() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for true {
			elem, err := lst.PopFront()
			if err != nil {
				break
			}
			if !yield(*elem) {
				return
			}
		}
	}
}
