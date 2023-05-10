package main

import "fmt"

func main(){
	a := 5
	defer printA(a)
	/*The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.*/
	a = 10
	fmt.Println("value of a before deferred function call", a)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}