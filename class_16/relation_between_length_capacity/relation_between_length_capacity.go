package main

import "fmt"

func main(){
	var mySlice []int
	//var mySlice = []int{}
	//mySlice := []int{}
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	y := make([]int, 5)
	fmt.Println(len(y))
	fmt.Println(cap(y))
}