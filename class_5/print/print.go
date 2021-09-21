/*Print formats using the default formats for its operands and writes to standard output. 
"Spaces are added between operands when neither is a string". It returns the number 
of bytes written and any write error encountered.*/

package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", "shihab"
	a, b:= 12, 23
	//fmt.Print(name, " is ", age, " years old.\n")
	fmt.Print(a,b,"\n")
	fmt.Print(name,age)

	// It is conventional not to worry about any
	// error returned by Print.

}