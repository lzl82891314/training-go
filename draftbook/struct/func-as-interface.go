package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Class struct{}

func (c *Class) Call(p interface{}) {
	fmt.Println(p, "from class")
}

type InvokeFunc func(interface{})

func (f InvokeFunc) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker Invoker
	clazz := &Class{}

	invoker = clazz

	invoker.Call("Hello world")

	invoker = InvokeFunc(func(i interface{}) {
		fmt.Println(i, "from function")
	})

	invoker.Call("Hello world")
}
