package main

import (
	"fmt"
)

func main() {
	var name, age = "Kim", 22
	s:= fmt.Sprintf("%s is %d years old.\n", name, age)

	fmt.Println(s)

}