package package1

import (
	"fmt"
	"github.com/lzl82891314/training-go/draftbook/struct/facotry-pattern/base"
)

type Instance01 struct {
}

func (clazz *Instance01) DoSomething() {
	fmt.Println("instance01 do something")
}

func init() {
	base.Register("Instance01", func() base.Instance {
		fmt.Println("instance02 init")
		return new(Instance01)
	})
}
