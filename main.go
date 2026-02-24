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
	bst := ds.NewBST[int]()

	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(15)
	bst.Insert(12)
	bst.Insert(20)

	bst.Remove(10)
	bst.Remove(7)
	bst.Remove(15)

	fmt.Println(bst.String())

	mx, _ := bst.Max()
	mn, _ := bst.Min()

	fmt.Println(*mx)
	fmt.Println(*mn)
}
