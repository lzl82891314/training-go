package draftbook

import "fmt"

func main1() {
	x := 1
	fmt.Println(x)
	add(&x)
	fmt.Println(x)

	y := make([]int, 2, 2)
	y[0] = 1
	y[1] = 2
	fmt.Println(y)
	addS(y)
	fmt.Println(y)

	z := [2]int{1, 2}
	fmt.Println(z)
	addA(z)
	fmt.Println(z)
}

func add(xp *int) {
	(*xp)++
}

func addS(s []int) {
	s[0] = 3
	s[1] = 4
}

func addA(a [2]int) {
	a[0] = 3
	a[1] = 4
}
