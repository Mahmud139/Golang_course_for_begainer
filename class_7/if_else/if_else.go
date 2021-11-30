package main

import "fmt"

func main(){
	num := 21
	if num%2 == 0 {
		fmt.Printf("The number %d is even\n", num)
	} else {
		fmt.Printf("THe number %d is odd\n", num)
	}
}