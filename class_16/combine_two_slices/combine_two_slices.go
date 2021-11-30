package main

import "fmt"

func main(){
	//slide no 28
	slice1 := []int{1,2,3}
	slice2 := []int{4,5,6}
	newSlice := append(slice1, slice2...)
	fmt.Println(newSlice)
}