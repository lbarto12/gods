package main

import (
	"fmt"
	"gods/ds"
)

func main() {
	stk := ds.NewStack[int]()

	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.Push(5)

	stk.Pop()
	fmt.Println(stk.String())

	stk.Pop()
	fmt.Println(stk.String())

	stk.Pop()
	fmt.Println(stk.String())

	stk.Pop()
	fmt.Println(stk.String())

	stk.Pop()
	fmt.Println(stk.String())

	stk.Pop()
	fmt.Println(stk.String())

}
