package main 

import "fmt"

func main(){
	var myBool bool = true
	fmt.Println("before calling printPointer(): ", myBool) 
	printPointer(&myBool)
	fmt.Println("after calling printPointer(): ", myBool)
}

func printPointer(myBoolPointer *bool) {
	*myBoolPointer = false
	fmt.Println("inside printPointer(): ",*myBoolPointer)
}