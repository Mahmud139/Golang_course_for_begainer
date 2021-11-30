package main

import "fmt"

func main(){
	num := 90

	if num <= 50 {
		fmt.Println(num, "is less than or equal to 50.")
	} else if num >= 50 && num <= 80 {
		fmt.Println(num, "is beetween 50 and 100.")
	} else if num >= 80 && num <= 100 {
		fmt.Println(num, "is beetween 80 and 100.")
	} else{
		fmt.Println(num, "is greater than 100.")
	}

}