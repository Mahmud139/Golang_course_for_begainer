// number conversion through division
package main

import (
	"fmt"
	"math"
)

func main(){
	// a:= 5 / 2
	// fmt.Println(a)

	// b:= 5 / 2.0
	// fmt.Println(b)

	// modulo 
	/* The % operator is the modulo, which returns the remainder rather than the
	quotient after division */
	// r := 20 % 3
	// fmt.Println(r)

	/* To do modulus math with float64 data types, you’ll use the Mod
	function from the math package: */

	m:= 12.121
	n:= 3.12
	res:= math.Mod(m,n)
	fmt.Println(res)
	
	
}