package main 

import "fmt"

func main(){
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	var myInt int
	fmt.Println(&myInt)
	var myFloat float64
	fmt.Println(&myFloat)
	var myBool bool
	fmt.Println(&myBool)
}