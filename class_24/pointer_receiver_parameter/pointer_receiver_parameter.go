package main

import "fmt"

type Number int

// func (n Number) Double() {
// 	n *= 2
// }

/*To make the double function work, we had to pass a pointer to the value we wanted to 
update, and then update the value at that pointer within the function.*/

func (n *Number) Double() {
	*n *= 2
}

func main() {
	number := Number(12) //variable with a nonpointer type
	fmt.Println("Original value of number:", number)
	number.Double() //Go will automatically convert the receiver to a pointer for you
	fmt.Println("number after calling double:", number)
}

/*When you call a method that requires a pointer receiver on a variable with a nonpointer 
type, Go will automatically convert the receiver to a pointer for you. The same is true 
for variables with pointer types; if you call a method requiring a value receiver, Go will 
automatically get the value at the pointer for you and pass that to the method.*/