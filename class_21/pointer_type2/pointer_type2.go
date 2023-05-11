package main

import "fmt"

func main(){
	var myInt int = 12
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer) //this print the address of myInt in HexCode
	fmt.Println(&myIntPointer) //this print the address of myIntPointer variable itself
	fmt.Println(*myIntPointer) //this print the value of myInt variable using the address.

	var myBool bool 
	// var myBoolPointer *bool 
	myBoolPointer := &myBool //short-hand declaration
	fmt.Println(myBoolPointer)
}