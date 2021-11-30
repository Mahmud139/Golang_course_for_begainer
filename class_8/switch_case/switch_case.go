package main

import "fmt"

func main(){
	/* Suppose we wanted to write a program that printed the English names for numbers.
	Using what we've learned so far we might start by doing this:*/
	i := 10

	// if i == 0 {
	// 	fmt.Println("zero")
	// } else if i == 1 {
	// 	fmt.Println("One")
	// } else if i == 2 {
	// 	fmt.Println("Two")
	// } else if i == 3 {
	// 	fmt.Println("Three")
	// } else if i == 4 {
	// 	fmt.Println("Four")
	// } else if i == 5 {
	// 	fmt.Println("Five")
	// } else {
	// 	fmt.Println("Invalid number.")
	// }

	/*Since writing a program in this way would be pretty tedious Go provides another 
	statement to make this easier: the switch statement. We can rewrite our program 
	to look like this:*/

	switch i {
	case 0: 
		fmt.Println("zero")
	case 1,6,7,8: 
		fmt.Println("One")
	case 2: 
		fmt.Println("Two")
	case 3: 
		fmt.Println("Three")
	case 4: 
		fmt.Println("Four")
	case 5: 
		fmt.Println("Five")
	default: 
		fmt.Println("Invalid number.")
	}



}