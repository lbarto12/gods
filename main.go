package main

import (
	"fmt"

	"github.com/lbarto12/gods/ds"
)

// func main() {
// 	q := ds.NewQueue[int]()

// 	q.Push(1)
// 	q.Push(2)
// 	q.Push(3)
// 	q.Push(4)

// 	for val, err := q.Pop(); err == nil; val, err = q.Pop() {
// 		fmt.Println(*val, " ", q.String())
// 	}

// }

func main() {

	lst := ds.NewList[int]()

	lst.PushFront(1)
	lst.PushFront(2)
	lst.PushFront(3)
	lst.PushFront(4)

	fmt.Println(lst.DumpSlice())

	fmt.Println(lst.String())
	for v := range lst.Iter() {
		fmt.Println(v)
	}

	fmt.Println(lst.String())

	lst.PopFront()
	lst.PopBack()
	lst.PushBack(2)

	fmt.Println(lst.Len())

}
