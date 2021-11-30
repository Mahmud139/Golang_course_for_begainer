package main

import "fmt"

func main(){
	num := 20

	if num >= 10 {
		if num >= 0 && num <= 15 {
			fmt.Println(num, "is between 0 to 15")
		} else {
			fmt.Println(num, "is greater than 15.")
		}
	} else {
		fmt.Println(num, "is less than 10.")
	}
}