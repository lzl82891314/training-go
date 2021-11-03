package slice

import (
	"fmt"
	"strings"
)

func main() {
	list := []string{"go hello", "go version", "go build", "go run"}

	chain := []func(string) string{
		removePrefixGo,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProcess(list, chain)

	for _, val := range list {
		fmt.Println(val)
	}
}

func removePrefixGo(str string) string {
	return strings.TrimPrefix(str, "go")
}

func StringProcess(list []string, chain []func(string) string) {
	for index, val := range list {
		result := val
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}
