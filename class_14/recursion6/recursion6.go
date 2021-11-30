//Tail recursion
package main

import "fmt"

func main(){
	printNum(3)
}

func printNum(n int) {
	if n > 0 {
		fmt.Println(n)
		printNum(n-1)
	}
}

/* this types of recursion are called Tail recursion, since the recursive call is 
the "last statement" in the function. there is no other statement or operation after 
the call. */