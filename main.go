package main

import (
	"fmt"

	"github.com/lbarto12/gods/ds/tree/bst"
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
	tree := bst.NewBST[int]()

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(15)
	tree.Insert(12)
	tree.Insert(20)

	tree.Remove(10)
	tree.Remove(7)
	tree.Remove(15)

	fmt.Println(tree.String())

	mx, _ := tree.Max()
	mn, _ := tree.Min()

	fmt.Println(*mx)
	fmt.Println(*mn)

	val, err := tree.Search(5)

	fmt.Println(*val, err)
}
