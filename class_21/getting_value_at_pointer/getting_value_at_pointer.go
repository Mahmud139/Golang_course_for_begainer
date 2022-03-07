package main

import "fmt"

func main(){
	myInt := 4 
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	myBool := true 
	myBoolPointer := &myBool 
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
}