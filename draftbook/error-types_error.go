package draftbook

import "fmt"

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"Something happened", "server.go", 42}
}

func main5() {
	err := test()
	switch err := err.(type) {
	case nil:
	case *MyError:
		fmt.Println("error occurred on line: ", err.Line)
	default:
	}
}
