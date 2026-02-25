package main

import (
	"fmt"

	"github.com/lbarto12/gods/ds/map/set"
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
	st1 := set.New(1, 2, 3)
	st2 := set.New(3, 4, 5)

	fmt.Println(st1.SymDiff(st2))
}
