package main

import "github.com/lzl82891314/training-go/draftbook/struct/facotry-pattern/base"
import _ "github.com/lzl82891314/training-go/draftbook/struct/facotry-pattern/package1"
import _ "github.com/lzl82891314/training-go/draftbook/struct/facotry-pattern/package2"

func main() {
	clazz01 := base.Create("Instance01")
	clazz01.DoSomething()

	clazz02 := base.Create("Instance02")
	clazz02.DoSomething()

	clazz03 := base.Create("Instance03")
	clazz03.DoSomething()
}
