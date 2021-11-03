package main

import (
	"fmt"
	"reflect"
)

type Container struct {
	s reflect.Value
}

func NewContainer(t reflect.Type, size int) *Container {
	if size < 0 {
		size = 64
	}
	return &Container{s: reflect.MakeSlice(reflect.SliceOf(t), 0, size)}
}

func (c *Container) Put(val interface{}) error {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("put error: cannot put a %T into a slice of %s", val, c.s.Type().Elem())
	}
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}

func (c *Container) Get(val interface{}) error {
	if reflect.ValueOf(val).Kind() != reflect.Ptr ||
		reflect.ValueOf(val).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("get error, needs *%s but got %T", c.s.Type().Elem(), val)
	}
	reflect.ValueOf(val).Elem().Set(c.s.Index(0))
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	container := NewContainer(reflect.TypeOf(nums[0]), 16)
	for _, n := range nums {
		if err := container.Put(n); err != nil {
			panic(err)
		}
	}

	num := -10
	if err := container.Get(0); err != nil {
		panic(err)
	}
	fmt.Println(num)
}
