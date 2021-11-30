package main

import "fmt"

func main(){
	grade:= 90

	// if grade >= 60 {
	// 	fmt.Println("You are passed!")
	// }
	// fmt.Println("if statement skipped")

	//return keyword
	// if grade >= 60 {
	// 	fmt.Println("You are passed!")
	// 	return 
	// }
	// fmt.Println("Study Hard")


	//tow conditions both are true or false
	if grade >= 60 && grade <= 70 {
		fmt.Println("your CGP will be low")
	}

	if 5 > 6 || 5 <= 6 {
		fmt.Println("it will be printed")
	}
}