//Passing functions as arguments to other functions
package main

import "fmt"

func main(){	
	simple(areaCal)

	// foo := func(a, b int) int {
	// 	return a * b
	// }
	// simple(foo)
	// a := 12
	// b := a 
	// b = 23
	// res :=areaCal(23,23) 
	// fmt.Println(res)
}

func areaCal(a, b int) int {
	return a * b 
}

//function simple takes a function that accepts two int arguments and returns a int as a parameter. 
func simple(area func(a, b int) int) {
	r := area(12,12)
	fmt.Println(r)
}




