package package2

import (
	"fmt"
	"github.com/lzl82891314/training-go/draftbook/struct/facotry-pattern/base"
)

type Instance02 struct {
}

func (clazz *Instance02) DoSomething() {
	fmt.Println("instance02 do something")
}

func init() {
	base.Register("Instance02", func() base.Instance {
		fmt.Println("instance02 init")
		return new(Instance02)
	})
}
