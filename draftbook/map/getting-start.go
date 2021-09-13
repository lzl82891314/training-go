package main

import "fmt"

func main() {
	var helloMap map[string]string
	helloMap = make(map[string]string, 10)
	helloMap["jeffery"] = "lizhenglong"
	helloMap["charis"] = "suihaochuan"
	helloMap["orren"] = "dengaoling"
	helloMap["curz"] = "liuzhonghe"

	fmt.Println(helloMap)
}
