package main

import "fmt"

func main(){
	r := factorial(5)
	fmt.Println(r)
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// please use the "recursion.png" picture to demonstrate how this function work

// func factorial(x int) int {
// 	if x == 0 || x == 1 {
// 		return 1
// 	}
	
// 	if x < 0 {
// 		fmt.Println("invalid number")
// 		return -1
// 	}

// 	return x * factorial(x-1)
// }


// these types of recursion are called direct recursion, since the function calls 
// itself directly, without the assistance of another function.