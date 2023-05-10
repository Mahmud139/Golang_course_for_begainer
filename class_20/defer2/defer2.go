package main

import "fmt"

func main(){
	fmt.Println(0)
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)


	// fmt.Println("Hello")
	// for i:= 1; i <= 3; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("World")
}