// indirect recursion 
package main

import "fmt"

func main(){
	firstFunc(2)

	//firstFunc(-1)
}

func firstFunc(n int) {
	if n >= 0 {
		fmt.Println("In first function:", n)
		secondFunc(n-1)
	}
}

func secondFunc(n int) {
	if n >= 0 {
		fmt.Println("In second function:", n)
		firstFunc(n-1)
	}
}

/* These types of recursion are called indirect recursion, since the function calls 
	another function and this function in turn, calls the calling function. */