package main

import (
	"fmt"
)

func main() {
	var name, age = "Kim", "shihab"
	a, b:= 12, 23
	//fmt.Print(name, " is ", age, " years old.\n")
	n:= fmt.Sprint(a,b)
	s:= fmt.Sprint(name,age)
	fmt.Println(n)
	fmt.Print(s)

	// It is conventional not to worry about any
	// error returned by Print.

}