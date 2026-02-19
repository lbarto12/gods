package main

import (
	"fmt"

	"github.com/lbarto12/gods/itertools"
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

	arr := []int{1, 2, 3, 4, 5}

	nw := itertools.Filter(arr, func(a int) bool {
		return a%2 == 0
	})

	fmt.Println(nw)

	arr2 := []int{6, 7, 8, 9, 10}
	arr3 := []int{12, 13, 14, 15, 16}

	zp := itertools.Zip(arr, arr2, arr3)

	fmt.Println(zp)

}
