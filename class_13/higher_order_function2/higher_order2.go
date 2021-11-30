//Returning functions from other functions
package main

import "fmt"

func main() {
	r := simple() //The return value from simple is assigned to r
	//Now r contains the function returned by simple function.
	shihab := r(12,14)
	fmt.Println(shihab)
	
	fmt.Println(r(12,12)) //We call r and pass it two int arguments

	// s := r(12,12)
	// fmt.Println(s)
}

func simple() func (a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f

	// return func(a, b int) int {
	// 	return a + b
	// }
}

