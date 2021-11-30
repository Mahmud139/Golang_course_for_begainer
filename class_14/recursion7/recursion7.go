// head recursion 
package main

import "fmt"

func main(){
	printNum(3)
}

func printNum(n int) {
	if n > 0 {
		printNum(n-1)
		fmt.Println(n)
	}
}

/* this types of recursion are called head recursion, since the recursive call is 
	the first statement in the function. there is no other statement or operation before 
	the call. */