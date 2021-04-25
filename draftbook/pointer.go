package draftbook

import "fmt"

type person struct {
	age  int
	name string
}

func (p *person) add() {
	p.age++
}

func addAge(p *person) {
	p.age++
}

func main3() {
	arr := [2]int{1, 2}
	p := &arr
	fmt.Println(p)
	fmt.Println(*p)
}
