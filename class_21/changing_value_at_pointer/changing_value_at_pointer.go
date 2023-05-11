package main

import "fmt"

func main(){
	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8  //myIntPointer == myInt are technically are the same thing
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}