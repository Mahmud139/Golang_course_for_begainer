/*Println formats using the default formats for its operands and writes to standard output. 
"Spaces are always added between operands and a newline is appended." It returns the number 
of bytes written and any write error encountered.*/

package main

import (
	"fmt"
)

func main() {
	//var name, age = "Kim", 24
	//n, m:= 12, 1212
	//fmt.Println(name, "is", age, "years old.")
	// fmt.Println(name,age)
	// fmt.Println(n,m)

	//prove that new line is appended ...we can use the return value of println function
	n, _ := fmt.Println("shihab")
	fmt.Print(n)

	// It is conventional not to worry about any
	// error returned by Println.

}
