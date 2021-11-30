package main

import (
	"fmt"
	// "reflect"
)

func main(){
	//slide no 27
	//create an array
	myArray := [5]int{1,2,3,4,5}
	//fmt.Println(reflect.TypeOf(myArray).Kind())
	//myArray = append(myArray, 23) // we can't append to myArray

	//converting an array to slice 
	mySlice := myArray[:]
	//fmt.Println(reflect.TypeOf(mySlice).Kind())
	mySlice[0] = 90
	// fmt.Println(mySlice)
	// fmt.Println(myArray)

	mySlice = append(mySlice, 23, 34)
	fmt.Println(mySlice, cap(mySlice))
	mySlice[4] = 9

	fmt.Println(myArray)
	fmt.Println(mySlice)
}