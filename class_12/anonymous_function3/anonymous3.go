//anonymous functions without assigning to a variable
package main

import "fmt"

func main(){
	//no parameter and no return value
	// func() { //an anonymous function is defined
	// 	fmt.Println("Hello world from a anonymous function.")
	// }() //after the function definition, we immediately call the function using paranthesis ()


	//with parameter and no return value
	// func(height int, width int) {
	// 	fmt.Println("Area is", height*width)
	// }(12,12)


	// with parameter and return value
	// func(height int, width int) int {
	// 	return height*width
	// }(12,12) 
	/*since this anonymous function have a return value, 
	so we need to store the value in a variable */
	// area := func(height int, width int) int {
	// 	return height*width
	// }(12,12) 
	// fmt.Println(area)


	//no parameter but have return value
	// func() string {
	// 	return "My name is Mahmudul Hasan."
	// }()
	//since this function return a value, so we need to store the value in a variable
	name := func() string {
		return "My name is Mahmudul Hasan."
	}()
	fmt.Println(name)

}