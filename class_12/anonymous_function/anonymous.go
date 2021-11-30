// anonymous function assigning to a variable
package main

import "fmt"

func main(){
	//without any parameter and return value
	a := func() { //we have assigned a function to the variable a
		fmt.Println("Hello world from a anonymous function.")
	}
	//This is the syntax for assigning a function to a variable.
	//If you notice carefully, the function assigned to "a" does not have a name. 
	//These kinds of functions are called anonymous functions since they do not have a name.
	
	a()
	b()
	fmt.Printf("The type of a() is: %T", a)
	//name := "shihab"

	//The only way to call this function is using the variable a. 
	//We have done this in the next line. a() calls the function 
	//and this prints "Hello world from a anonymous function."
}

var b = func ()  {
	fmt.Println("my name is shihab")
}

