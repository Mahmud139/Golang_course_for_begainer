//simple closure function. 
package main

import "fmt"

func main(){
	//var name string = "shihab"
	// add := func(a, b int) int {
	// 	return a + b
	// }
	// fmt.Println(add(12,12))

	/*When you create a local function like this it also
	has access to other local variables.*/
	x := 0
	increment := func() int {
		x++ //the anonymous function accesses the variable x which is present outside its body
		return x
	}
	fmt.Println(increment()) 
	/*increment adds 1 to the variable x which is defined in
	the main function's scope.*/
	fmt.Println(increment())
	fmt.Println(increment())
	
	/*This x variable can be accessed and modified by the increment function.*/
	/*A function like this together with the non-local variables it references 
	is known as a closure.*/
}