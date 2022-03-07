package main 

import "fmt"

func main(){
	var myBool bool = true 
	printPointer(&myBool)
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}