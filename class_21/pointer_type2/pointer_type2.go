package main

import "fmt"

func main(){
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myBool bool 
	var myBoolPointer *bool 
	myBoolPointer = &myBool
	fmt.Println(myBoolPointer)
}