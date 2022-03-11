package main

import "fmt"

func main() {
	var myStruct struct {
		name   string
		age    int
		status bool
	}

	myStruct.name = "Mahmud"
	myStruct.age = 23
	myStruct.status = true

	fmt.Println(myStruct.name)
	fmt.Println(myStruct.age)
	fmt.Println(myStruct.status)
	fmt.Println(myStruct)
}
