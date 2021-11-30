// anonymous function recursion 
package main

import "fmt"

/* Recursion can also be carried out using anonymous functions in Golang as shown below:*/

func main() {
	var recursiveAnonymous func(int)
	recursiveAnonymous = func(v int) {
		if v == -1 {
			fmt.Println("Hello, How are you?")
			return
		} else {
			fmt.Println(v)
			recursiveAnonymous(v-1)
		}
	}

	recursiveAnonymous(3)
}

