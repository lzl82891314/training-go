package main

import "fmt"

func main() {
	type MyInt = int64
	var data MyInt = 199
	fmt.Printf("%T", data)

	staffs := initStaffs()
	fmt.Println(staffs)
}

func initStaffs() map[string]*Staff {
	staffs := make(map[string]*Staff)
	staffs["jeffery"] = &Staff{name: "jeffery", age: 28, salary: 1000}
	staffs["chairs"] = &Staff{name: "chairs", age: 29, salary: 10000}
	staffs["orren"] = &Staff{name: "orren", age: 27, salary: 15000}
	staffs["alex"] = &Staff{name: "alex", age: 22, salary: 8000}
	return staffs
}

type Staff struct {
	name   string
	age    int
	salary int
}
