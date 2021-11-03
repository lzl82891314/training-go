package slice

import (
	"fmt"
	"math"
)

func main2() {
	capacity := 4
	var a = make([]int, 3, capacity) // capacity=3 和capacity=4会得到截然不同的输出
	a[0] = 1
	a[1] = 2
	a[2] = 3

	var b = a[:]
	a = append(a, 4) // append的时候还需要赋值一下. 语言设计者该为此感到耻辱吧
	a[0] = 1000

	fmt.Println(a)
	fmt.Println(b)
	c := math.MaxInt32
	fmt.Println(c)
}

// 实现删除
func RemoveLast(a []int) []int {
	length := len(a)
	a = (a)[:length-1]
	return a
}

func RemoveIndex(a []int, index int) []int {
	a = append(a[:index], a[index+1:]...)
	return a
}

func RemoveIndexP(a *[]int, index int) *[]int {
	*a = append((*a)[:index], (*a)[index+1:]...)
	return a
}
