package main

import "fmt"

func main(){
	num := 10

	if num >= 12 {
		fmt.Println(num, "is greater than 12.")
	} else{
		if num == 12 {
			fmt.Println(num, "is equal to 12.")
		} else{
			fmt.Println(num, "is not equal to 12.")
		}
	}
}