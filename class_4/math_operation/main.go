// math operation example
package main

import (
	"fmt"
)

func main(){
	// addition:= 1+2
	// subtraction:= 2-1
	// multiplication:= 2*3
	// division:= 6/3
	// remainder:= 6%2
	// name:="mahmudul"
	// name2:= name + "hasan"
	/* Concatenation uses the same symbol as addition. The Go compiler figures out 
	what to do based on the types of the arguments. Since both sides of the + are 
	strings the compiler assumes you mean concatenation and not addition. 
	(Addition is meaningless for strings)  */
	
	//fmt.Println(addition, subtraction, multiplication, division, remainder, name2)

	var n string ="shihab "
	var m string = "shihab"
	fmt.Println(n == m)
	fmt.Println(n <= m)

	
}