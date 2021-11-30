package main

import (
	"fmt"
)

func main() {
	//var name, age = "Kim", "shihab"
	//n, m:= 12, 1212
	//s := fmt.Sprintln(name, "is", age, "years old.")
	//s := fmt.Sprintln(name,age)
	//s := fmt.Sprintln(n,m)
	//fmt.Println(s)

	//prove that new line is appended ...we can use the return value of println function
	s:= fmt.Sprintln("shihab","uddin")
	fmt.Println(s)
	fmt.Println(len(s))

	// It is conventional not to worry about any
	// error returned by Println.

}