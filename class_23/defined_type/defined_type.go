package main

import "fmt"

func main(){
	var subcriber1 struct {
		name string
		rate float64
		active bool 
	}
	subcriber1.name = "Mahmud"
	fmt.Println(subcriber1.name)

	var subcriber2 struct {
		name string
		rate float64
		active bool 
	}
	subcriber2.name = "Hasan"
	fmt.Println(subcriber2.name)

	/* Declaring struct variable is really tedious for us. we have ro repeat entire struct type
	declaration for each new variable. */
}