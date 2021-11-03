package main

import (
	"fmt"
	"github.com/lzl82891314/training-go/draftbook/slice"
)

func main() {
	a1 := make([]int, 5, 10)
	a1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := a1
	//fmt.Println(a1)

	a1 = slice.RemoveIndex(a1, 5)
	fmt.Println(a1)
	fmt.Println(a2)

	p := slice.RemoveIndexP(&a1, 5)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(*p)
}
